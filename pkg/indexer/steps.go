package indexer

import (
	"context"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
)

type NextFunc = func(nextStep database.CreationFlow)

type Step = func(ctx context.Context, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink)

type StepMap = map[database.CreationFlow]Step

var Steps StepMap = map[database.CreationFlow]Step{
	database.CollectionCreator: FindCollectionCreator,
}
