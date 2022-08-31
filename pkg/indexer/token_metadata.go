package indexer

import (
	"context"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	reservoir "nft-indexer/pkg/indexer/thirdparty/reservoir"
)

func GetTokenMetadata(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	// totalSupply := collection.ZoraStats.NumNfts
	cursor := collection.State.Create.ReservoirCursor

	reservoir, err := reservoir.NewApiClient(config.Reservoir.ApiKey)
	if err != nil {
		panic(err)
		return
	}

	for true {
		res, err := reservoir.GetTokensInfo(collection.Address, 50, cursor)
		if err != nil {
			sink.WriteError(err, database.CollectionMetadata)
			return
		}

		for _, tokenWrapper := range res.Tokens {
			token := tokenWrapper.Token

			if token.TokenId != "" {
				// TODO: store token info
			}
		}

		if res.Continuation != "" {
			collection.State.Create.ReservoirCursor = res.Continuation
		} else {
			break
		}
	}
}
