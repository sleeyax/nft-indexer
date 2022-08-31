package indexer

import (
	"context"
	"fmt"
	"log"
	"math"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"nft-indexer/pkg/indexer/thirdparty/zora"
	"nft-indexer/pkg/utils"
	"strconv"
	"time"
)

func GetTokensAndMintInfo(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	totalSupply := collection.ZoraStats.NumNfts
	limit := 500

	z, err := zora.NewGraphQLClient(config.Zora.ApiKey)
	if err != nil {
		panic(err)
		return
	}

	var pageNr int

	for true {
		pageNr++

		res, err := z.GetTokens(collection.Address, limit, collection.State.Create.ZoraCursor)
		if err != nil {
			sink.WriteError(err, database.CollectionMintsStep)
			return
		}

		for _, tokenWrapper := range res.Data.Tokens.Nodes {
			token := tokenWrapper.Token

			if token.TokenId == "" {
				continue
			}

			metadata := database.Erc721TokenMetadata{
				Name:        token.Name,
				Title:       token.Name,
				Image:       token.Image.Url,
				Description: token.Description,
			}

			for _, attr := range token.Attributes {
				metadata.Attributes = append(metadata.Attributes, database.Erc721TokenAttribute{
					Value:     attr.Value,
					TraitType: attr.TraitType,
				})
			}

			var mintedAt int64 = 0
			if token.MintInfo.MintContext.BlockTimestamp != "" {
				if t, err := time.Parse(time.RFC3339, token.MintInfo.MintContext.BlockTimestamp); err != nil {
					mintedAt = t.UnixMilli()
				}
			}

			erc721token := database.Erc721Token{
				Slug:                 utils.ToSearchFriendly(token.Name),
				TokenId:              token.TokenId,
				ChainId:              collection.ChainId,
				CollectionAddress:    collection.Address,
				NumTraitTypes:        len(token.Attributes),
				Metadata:             metadata,
				UpdatedAt:            time.Now().UnixMilli(),
				Owner:                token.Owner,
				TokenStandard:        database.ERC721,
				MintedAt:             mintedAt,
				Minter:               database.Normalize(token.MintInfo.OriginatorAddress),
				MintTxHash:           token.MintInfo.MintContext.TransactionHash,
				MintPrice:            token.MintInfo.Price.ChainTokenPrice.Decimal,
				MintCurrencyAddress:  token.MintInfo.Price.ChainTokenPrice.Currency.Address,
				MintCurrencyDecimals: token.MintInfo.Price.ChainTokenPrice.Currency.Decimals,
				MintCurrencyName:     token.MintInfo.Price.ChainTokenPrice.Currency.Name,
			}

			if tokenIdNumeric, err := strconv.Atoi(token.TokenId); err == nil {
				erc721token.TokenIdNumeric = tokenIdNumeric
			}

			if token.Image.Url != "" {
				erc721token.Image = database.Erc721TokenImage{
					Url:       token.Image.Url,
					UpdatedAt: time.Now().UnixMilli(),
				}
			}

			collection.Tokens = append(collection.Tokens, erc721token)
			collection.NumNfts++
		}

		percentage := int(math.Floor(((float64(pageNr)*float64(limit))/float64(totalSupply))*100.0*100.0) / 100.0)

		// TODO: write progress to db/sink per iteration
		collection.State.Create = database.Create{
			Progress:  percentage,
			Step:      database.AggregateMetadataStep,
			UpdatedAt: time.Now().Unix(),
		}

		log.Println(fmt.Sprintf("fetching tokens and mints of %s:%s (%d %%)", collection.ChainId, collection.Address, percentage))

		if res.Data.Tokens.PageInfo.HasNextPage {
			collection.State.Create.ZoraCursor = res.Data.Tokens.PageInfo.EndCursor
		} else {
			break
		}
	}

	sink.Write(IndexResult{Collection: collection, Step: database.CollectionMintsStep})
}
