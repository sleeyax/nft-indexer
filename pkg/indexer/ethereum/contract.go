package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"nft-indexer/pkg/indexer/ethereum/tokens"
)

type Contract struct {
	Address   string
	NetworkId Network
	provider  *Provider
}

// NewContract creates a new Ethereum contract.
func NewContract(address string, networkId Network, provider *Provider) *Contract {
	return &Contract{
		Address:   address,
		NetworkId: networkId,
		provider:  provider,
	}
}

// GetHexAddress returns the hex address of the contract.
func (c *Contract) GetHexAddress() common.Address {
	return common.HexToAddress(c.Address)
}

// ToErc721 creates a new instance of tokens.Erc721, bound to this contract.
func (c *Contract) ToErc721() (*tokens.Erc721, error) {
	return tokens.NewErc721(c.GetHexAddress(), c.provider.ethereum)
}

// ToOwnable creates a new instance of tokens.Ownable, bound to this contract.
func (c *Contract) ToOwnable() (*tokens.Ownable, error) {
	return tokens.NewOwnable(c.GetHexAddress(), c.provider.ethereum)
}
