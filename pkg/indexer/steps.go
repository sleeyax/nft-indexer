package indexer

import (
	"context"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"time"
)

type NextFunc = func(nextStep database.CreationFlow)

type Step = func(ctx context.Context, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink)

type StepMap = map[database.CreationFlow]Step

var Steps StepMap = map[database.CreationFlow]Step{
	database.Unindexed: func(ctx context.Context, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
		// first step resets the collection
		collection.IndexInitiator = database.Normalize(ethereum.NullAddress.String())
		collection.ChainId = string(tokenContract.Contract().NetworkId)
		collection.Address = database.Normalize(tokenContract.Contract().Address.String())
		collection.TokenStandard = tokenContract.TokenStandard
		collection.HasBlueCheck = false
		collection.State = database.State{
			Version: 1,
			Create: database.Create{
				Progress:  0,
				Step:      database.CollectionCreator, // TODO: the main indexer func should be able to read the next step (based on the current one) automatically
				UpdatedAt: time.Now().Unix(),
			},
			Export: database.Export{Done: false},
		}

		sink.Write(IndexResult{
			Collection: collection,
			Step:       database.Unindexed,
		})
	},
	database.CollectionCreator: FindCollectionCreator,
}
