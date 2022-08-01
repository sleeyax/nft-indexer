package indexer

import (
	"context"
	"fmt"
	"io"
	"nft-indexer/pkg/config"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
)

type Indexer struct {
	io.Closer
	config   *config.Configuration
	chain    Chain
	provider *ethereum.Provider // TODO: this should become an interface if we ever want to support multiple chains
}

type IndexResult struct {
	// Currently accumulated NFT collection data.
	Collection *database.NFTCollection

	// Fatal error message.
	// If this field is set, the function that handles the Step has stopped processing information.
	Error error

	// Nonfatal error message.
	// Unlike Error, if this field is set, the function that handles the Step will continue processing information.
	Warning error

	// Current indexing step.
	Step database.CreationFlow
}

// New creates a new NFT indexer.
//
// The third argument can contain any parameters that are required to initialize a connection to the chain.
func New(config *config.Configuration, chain Chain, chainConfig interface{}) (*Indexer, error) {
	i := &Indexer{config: config, chain: chain}

	switch chain {
	case Ethereum:
		network := chainConfig.(ethereum.Network)

		// create and connect ethereum provider
		provider := ethereum.NewProvider(config)
		if err := provider.Connect(network); err != nil {
			return nil, err
		}

		i.provider = provider
	default:
		return nil, fmt.Errorf("chain %s is currently unsupported", chain)
	}

	return i, nil
}

// Start starts the indexer flow.
func (i *Indexer) Start(ctx context.Context, collection *database.NFTCollection, ch chan IndexResult) {
	defer close(ch)

	sink := NewSink(ch)

	if i.chain != Ethereum {
		sink.WriteError(fmt.Errorf("chain %s is currently unsupported", i.chain))
		return
	}

	// create contract
	contract := ethereum.NewContract(collection.Address, ethereum.Network(collection.ChainId), i.provider)

	// parse the contract into an ERC-721 compatible token contract
	tokenContract, err := ethereum.NewTokenContract(contract, database.ERC721)
	if err != nil {
		sink.WriteError(err)
		return
	}

	if collection.State.Create.Step == "" {
		collection.State.Create.Step = database.CollectionCreator
	}

	for {
		select {
		default:
			if collection.State.Create.Step == database.Complete {
				return
			}

			stepFunc, ok := Steps[collection.State.Create.Step]
			if !ok {
				return
			}

			stepFunc(ctx, tokenContract, collection, sink)
		case <-ctx.Done():
			return
		}
	}
}

func (i *Indexer) Close() error {
	return i.provider.Close()
}
