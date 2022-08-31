package indexer

import (
	"context"
	"fmt"
	"log"
	"math"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	reservoir "nft-indexer/pkg/indexer/thirdparty/reservoir"
	"nft-indexer/pkg/utils"
	"strconv"
	"time"
)

func GetTokenMetadata(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	totalSupply := collection.ZoraStats.NumNfts
	limit := 50

	rv, err := reservoir.NewApiClient(config.Reservoir.ApiKey)
	if err != nil {
		panic(err)
		return
	}

	var pageNr int

	for true {
		pageNr++

		res, err := rv.GetTokensInfo(collection.Address, limit, collection.State.Create.ReservoirCursor)
		if err != nil {
			sink.WriteError(err, database.CollectionMetadataStep)
			return
		}

		for _, tokenWrapper := range res.Tokens {
			token := tokenWrapper.Token

			if token.TokenId == "" {
				continue
			}

			metadata := database.Erc721TokenMetadata{
				Name:        token.Name,
				Title:       token.Name,
				Image:       token.Image,
				Description: token.Description,
			}

			for _, attr := range token.Attributes {
				metadata.Attributes = append(metadata.Attributes, database.Erc721TokenAttribute{
					Value:     attr.Value,
					TraitType: attr.Key,
				})
			}

			erc721token := database.Erc721Token{
				Slug:              utils.ToSearchFriendly(token.Name),
				TokenId:           token.TokenId,
				ChainId:           collection.ChainId,
				CollectionAddress: collection.Address,
				NumTraitTypes:     len(token.Attributes),
				Metadata:          metadata,
				UpdatedAt:         time.Now().UnixMilli(),
				Owner:             token.Owner,
				TokenStandard:     database.ERC721,
			}

			if tokenIdNumeric, err := strconv.Atoi(token.TokenId); err == nil {
				erc721token.TokenIdNumeric = tokenIdNumeric
			}

			if token.Image != "" {
				erc721token.Image = database.Erc721TokenImage{
					Url:       token.Image,
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
			Step:      database.CollectionMintsStep,
			UpdatedAt: time.Now().Unix(),
		}

		log.Println(fmt.Sprintf("fetching tokens of %s:%s (%d %%)", collection.ChainId, collection.Address, percentage))

		if res.Continuation != "" {
			collection.State.Create.ReservoirCursor = res.Continuation
		} else {
			break
		}
	}

	sink.Write(IndexResult{Collection: collection, Step: database.TokenMetadataStep})
}
