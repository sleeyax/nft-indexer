package indexer

import (
	"context"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"time"
)

func FindCollectionCreator(ctx context.Context, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	step := database.CollectionCreator

	ownableContract, err := tokenContract.ToOwnable()
	if err != nil {
		sink.WriteError(err, step)
		return
	}

	// read contract creation event data
	creationEvent, err := ownableContract.GetCreationEvent()
	if err != nil {
		sink.WriteError(err, step)
		return
	}

	collection.Deployer = database.Normalize(creationEvent.NewOwner.String())
	collection.DeployedAtBlock = int(creationEvent.Raw.BlockNumber)

	// read block
	deployedAt, err := tokenContract.Contract().ReadTimestamp(ctx, creationEvent.Raw.BlockHash)
	if err != nil {
		sink.WriteWarning(err, step)
	} else {
		collection.DeployedAt = int(deployedAt * 1000)
	}

	// try to find the owner or creator
	var owner string
	owner, err = ownableContract.GetOwner()
	if err != nil {
		sink.WriteWarning(err, step)
	}
	if owner == ethereum.NullAddress.String() {
		owner = creationEvent.NewOwner.String()
	}

	collection.Owner = database.Normalize(owner)

	collection.State.Create = database.Create{
		Step:      database.CollectionMetadata,
		UpdatedAt: time.Now().Unix(),
	}

	sink.Write(IndexResult{Collection: collection, Step: step})
}
