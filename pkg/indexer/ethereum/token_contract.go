package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"nft-indexer/pkg/database"
)

var nullAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

type TokenContract struct {
	contract      *Contract
	tokenStandard database.TokenStandard
}

func NewTokenContract(contract *Contract, standard database.TokenStandard) (*TokenContract, error) {
	return &TokenContract{
		contract,
		standard,
	}, nil
}

// GetCreator returns the address of the person who initially created the token contract.
// Requires OpenZeppelin's Ownable.sol to be implemented by the contract.
func (tc *TokenContract) GetCreator() (string, error) {
	ownable, err := tc.contract.ToOwnable()
	if err != nil {
		return "", err
	}

	iterator, err := ownable.FilterOwnershipTransferred(nil, []common.Address{nullAddress}, []common.Address{})
	if err != nil {
		return "", err
	}

	if !iterator.Next() {
		return "", fmt.Errorf("failed to find the creator contract %s on network %s", tc.contract.Address, tc.contract.NetworkId)
	}

	return iterator.Event.NewOwner.Hex(), nil
}

// TODO: get current owner of the contract
