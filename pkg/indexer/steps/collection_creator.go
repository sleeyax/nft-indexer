package steps

import (
	"context"
	"nft-indexer/pkg/database"
)

// CollectionCreatorStep is a step to find the collection creator.
type CollectionCreatorStep struct{}

func (c CollectionCreatorStep) Execute(context context.Context, collection *database.NFTCollection) (database.CreationFlow, error) {
	//TODO implement me
	panic("implement me")
}
