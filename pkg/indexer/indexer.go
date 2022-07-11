package indexer

import (
	"context"
	"errors"
	"fmt"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
)

type Indexer struct {
	config *nft_indexer.Configuration
}

func New(config *nft_indexer.Configuration) *Indexer {
	return &Indexer{config: config}
}

// Start starts the indexer flow.
func (i *Indexer) Start(ctx context.Context, collection *database.NFTCollection, steps StepMap) error {
	if collection.State.Create.Step == "" {
		collection.State.Create.Step = database.CollectionCreator
	}

	currentStep := collection.State.Create.Step

	for {
		step, ok := steps[currentStep]
		if !ok {
			return errors.New(fmt.Sprintf("Step '%s' is not yet implemented!\n", collection.State.Create.Step))
		}

		nextStep, err := step.Execute(ctx, collection)
		if err != nil {
			return err
		}

		if currentStep == database.Complete {
			break
		}

		currentStep = nextStep
	}

	return nil
}
