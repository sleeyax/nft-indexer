package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"nft-indexer/pkg/indexer/ethereum/tokens"
)

type Contract struct {
	Address   common.Address
	NetworkId Network
	provider  *Provider
}

// NewContract creates a new Ethereum contract.
func NewContract(address string, networkId Network, provider *Provider) *Contract {
	return &Contract{
		Address:   common.HexToAddress(address),
		NetworkId: networkId,
		provider:  provider,
	}
}

// ToErc721 creates a new instance of tokens.Erc721, bound to this contract.
func (c *Contract) ToErc721() (*tokens.Erc721, error) {
	return tokens.NewErc721(c.Address, c.provider.ethereum)
}

// ToOwnable creates a new instance of tokens.Ownable, bound to this contract.
func (c *Contract) ToOwnable() (*tokens.Ownable, error) {
	return tokens.NewOwnable(c.Address, c.provider.ethereum)
}
