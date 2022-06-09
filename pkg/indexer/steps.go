package indexer

import (
	"nft-indexer/pkg/database"
	steps2 "nft-indexer/pkg/indexer/steps"
)

// DefaultSteps is a default set of steps the indexer should perform from start to finish in order to deem an NFT collection as complete.
var DefaultSteps StepMap

func init() {
	DefaultSteps = map[database.CreationFlow]Step{
		database.CollectionCreator: new(steps2.CollectionCreatorStep), // TODO: just use a function (as part of a struct?) anyways?
	}
}
