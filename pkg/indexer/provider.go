package indexer

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	nft_indexer "nft-indexer"
)

// EthereumNetwork indicates the type of network to connect to (main or test).
type EthereumNetwork string

const (
	MainNetwork   EthereumNetwork = "1"
	GoerliNetwork EthereumNetwork = "5"
)

type Provider struct {
	config   *nft_indexer.Configuration
	ethereum *ethclient.Client
}

func NewProvider(config *nft_indexer.Configuration) *Provider {
	return &Provider{config: config}
}

// ConnectEthereum connects to the specified ethereum network over JSON RPC.
func (p *Provider) ConnectEthereum(network EthereumNetwork) error {
	switch network {
	case GoerliNetwork:
	case MainNetwork:
		var rpcUrls []string
		if network == MainNetwork {
			rpcUrls = p.config.Alchemy.JsonRpc.MainNet
		} else {
			rpcUrls = p.config.Alchemy.JsonRpc.Goerli
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
