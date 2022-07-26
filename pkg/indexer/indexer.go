package indexer

import (
	"context"
	"fmt"
	"io"
	"log"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
	"time"
)

type Indexer struct {
	io.Closer
	config   *nft_indexer.Configuration
	chain    Chain
	provider *ethereum.Provider // TODO: this should become an interface if we ever want to support multiple chains
}

type IndexResult struct {
	Collection *database.NFTCollection
	Error      error
}

// New creates a new NFT indexer.
//
// The third argument can contain any parameters that are required to initialize a connection to the chain.
func New(config *nft_indexer.Configuration, chain Chain, chainConfig interface{}) (*Indexer, error) {
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

	if i.chain != Ethereum {
		ch <- IndexResult{
			Error: fmt.Errorf("chain %s is currently unsupported", i.chain),
		}
		return
	}

	// create contract
	contract := ethereum.NewContract(collection.Address, ethereum.Network(collection.ChainId), i.provider)

	// parse the contract into an ERC-721 compatible tokenContract contract
	tokenContract, err := ethereum.NewTokenContract(contract, database.ERC721)
	if err != nil {
		ch <- IndexResult{Error: err}
		return
	}

	if collection.State.Create.Step == "" {
		collection.State.Create.Step = database.CollectionCreator
	}

	for {
		select {
		default:
			switch collection.State.Create.Step {
			case database.CollectionCreator:
				// first step resets the colleciton
				collection.IndexInitiator = ethereum.NullAddress.String()
				collection.ChainId = string(contract.NetworkId)
				collection.Address = contract.Address.String()
				collection.TokenStandard = tokenContract.TokenStandard
				collection.HasBlueCheck = false

				// try to find the owner or creator
				var owner string
				owner, err = tokenContract.GetOwner()
				if owner == ethereum.NullAddress.String() {
					owner, err = tokenContract.GetCreator()
				}
				if err != nil {
					log.Printf("failed to find collection owner or creator: %e", err)
				}

				collection.Owner = owner

				collection.State.Create = database.Create{
					Progress:  0,
					Step:      database.CollectionMetadata,
					UpdatedAt: time.Now().Unix(),
				}
				collection.State.Version = 1
				collection.State.Export = database.Export{Done: collection.State.Export.Done}

				ch <- IndexResult{Collection: collection}
			default:
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (i *Indexer) Close() error {
	return i.provider.Close()
}
