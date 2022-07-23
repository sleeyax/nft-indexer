package indexer

import (
	"context"
	"fmt"
	"log"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/database"
	"nft-indexer/pkg/indexer/ethereum"
)

type Indexer struct {
	config   *nft_indexer.Configuration
	chain    Chain
	provider *ethereum.Provider // TODO: this should become an interface if we ever want to support multiple chains
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
func (i *Indexer) Start(ctx context.Context, collection *database.NFTCollection) error {
	if i.chain != Ethereum {
		return fmt.Errorf("chain %s is currently unsupported", i.chain)
	}

	defer i.provider.Close()

	// create contract
	contract := ethereum.NewContract(collection.Address, ethereum.Network(collection.ChainId), i.provider)

	// parse the contract into an ERC-721 compatible tokenContract contract
	tokenContract, err := ethereum.NewTokenContract(contract, database.ERC721)
	if err != nil {
		return err
	}

	if collection.State.Create.Step == "" {
		collection.State.Create.Step = database.CollectionCreator
	}

	for {
		switch collection.State.Create.Step {
		case database.CollectionCreator:
			// find owner or creator
			var owner string
			owner, err = tokenContract.GetOwner()
			if err == ethereum.NullAddressError {
				owner, err = tokenContract.GetCreator()
			}

			if err != nil {
				return err
			}

			log.Println(owner)
			return nil // TODO: write to DB
		}
	}
}
