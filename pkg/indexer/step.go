package indexer

import (
	"context"
	"nft-indexer/pkg/database"
)

type StepMap map[database.CreationFlow]Step

type Step interface {
	// Execute executes this step.
	// The specified collection may be modified during this step and execution may be canceled via the provided context.
	Execute(context context.Context, collection *database.NFTCollection) (database.CreationFlow, error)
}
