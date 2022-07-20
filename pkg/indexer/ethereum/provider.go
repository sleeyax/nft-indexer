package ethereum

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	nft_indexer "nft-indexer"
	"nft-indexer/pkg/indexer"
)

type Provider struct {
	indexer.Provider
	ethereum *ethclient.Client
}

// NewProvider creates a new ethereum provider instance.
func NewProvider(config *nft_indexer.Configuration) *Provider {
	return &Provider{
		Provider: indexer.Provider{Config: config},
	}
}

// Connect connects to the specified ethereum network over JSON RPC.
func (p *Provider) Connect(network Network) error {
	switch network {
	case GoerliNetwork:
	case MainNetwork:
		var rpcUrls []string
		if network == MainNetwork {
			rpcUrls = p.Config.Alchemy.JsonRpc.MainNet
		} else {
			rpcUrls = p.Config.Alchemy.JsonRpc.Goerli
		}

		rpcUrl := nft_indexer.RandomItem(rpcUrls)

		conn, err := ethclient.Dial(rpcUrl)
		if err != nil {
			return err
		}

		p.ethereum = conn

		return nil
	}

	return errors.New(fmt.Sprintf("network '%s' not found", network))
}

func (p *Provider) Close() error {
	p.ethereum.Close()
	return nil
}
