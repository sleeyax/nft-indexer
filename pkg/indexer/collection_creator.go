package indexer

import (
	"context"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"time"
)

func FindCollectionCreator(ctx context.Context, tokenContract *ethereum.TokenContract, collection *database.NFTCollection, sink *Sink) {
	step := database.CollectionCreator

	// first step resets the collection
	// TODO: move this to different step
	collection.IndexInitiator = database.Normalize(ethereum.NullAddress.String())
	collection.ChainId = string(tokenContract.Contract().NetworkId)
	collection.Address = database.Normalize(tokenContract.Contract().Address.String())
	collection.TokenStandard = tokenContract.TokenStandard
	collection.HasBlueCheck = false
	collection.State.Version = 1
	collection.State.Export = database.Export{Done: false}

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
		Progress:  0,
		Step:      database.CollectionMetadata,
		UpdatedAt: time.Now().Unix(),
	}

	sink.Write(IndexResult{Collection: collection, Step: step})
}
