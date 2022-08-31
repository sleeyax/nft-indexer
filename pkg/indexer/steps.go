package indexer

import (
	"context"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"time"
)

type Step = func(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink)

type StepMap = map[database.CreationStep]Step

var Steps StepMap = map[database.CreationStep]Step{
	database.UnindexedStep: func(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
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
				Step:      database.CollectionCreatorStep, // TODO: the main indexer func should be able to read the next step (based on the current one) automatically
				UpdatedAt: time.Now().Unix(),
			},
			Export: database.Export{Done: false},
		}

		sink.Write(IndexResult{
			Collection: collection,
			Step:       database.UnindexedStep,
		})
	},
	database.CollectionCreatorStep:  FindCollectionCreator,
	database.CollectionMetadataStep: GetCollectionMetadata,
	// NOTE: this step is skipped because the data is already retrieved in the next step
	// see: https://github.com/sleeyax/nft-indexer/issues/12
	// database.TokenMetadataStep:      GetTokenMetadata,
	database.TokenMetadataStep: func(ctx context.Context, config *config.Configuration, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
		collection.State = database.State{
			Create: database.Create{
				Step:      database.CollectionMintsStep,
				UpdatedAt: time.Now().Unix(),
			},
		}

		sink.Write(IndexResult{
			Collection: collection,
			Step:       database.TokenMetadataStep,
		})
	},
	database.CollectionMintsStep: GetTokensAndMintInfo,
}
