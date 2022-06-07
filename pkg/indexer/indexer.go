package indexer

import (
	"context"
	"errors"
	"fmt"
	"nft-indexer/pkg/database"
)

func Index(ctx context.Context, collection *database.NFTCollection, steps StepMap) error {
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
