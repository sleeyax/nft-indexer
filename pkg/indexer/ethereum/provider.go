package ethereum

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	config "nft-indexer/pkg/config"
	"nft-indexer/pkg/utils"
)

type Provider struct {
	config   *config.Configuration
	ethereum *ethclient.Client
}

// NewProvider creates a new ethereum provider instance.
func NewProvider(cfg *config.Configuration) *Provider {
	return &Provider{config: cfg}
}

// Connect connects to the specified ethereum network over JSON RPC.
func (p *Provider) Connect(network Network) error {
	switch network {
	case GoerliNetwork:
	case MainNetwork:
		var rpcUrls []string
		if network == MainNetwork {
			rpcUrls = p.config.Alchemy.JsonRpc.MainNet
		} else {
			rpcUrls = p.config.Alchemy.JsonRpc.Goerli
		}

		rpcUrl := utils.RandomItem(rpcUrls)

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
