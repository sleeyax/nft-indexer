package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"nft-indexer/pkg/indexer/ethereum/tokens"
)

// OwnableTokenContract is a TokenContract that implements OpenZeppelin's Ownable.sol contract.
type OwnableTokenContract struct {
	*TokenContract
	ownable *tokens.Ownable
}

// GetCreator returns the address of the person who initially created the token contract.
func (otc *OwnableTokenContract) GetCreator() (string, error) {
	iterator, err := otc.ownable.FilterOwnershipTransferred(nil, []common.Address{NullAddress}, []common.Address{})
	if err != nil {
		return "", err
	}

	if !iterator.Next() {
		return "", fmt.Errorf("failed to find the creator contract %s on network %s", otc.contract.Address, otc.contract.NetworkId)
	}

	return iterator.Event.NewOwner.Hex(), nil
}

// GetOwner returns the address of the contract owner, which is stored in the contract.
func (otc *OwnableTokenContract) GetOwner() (string, error) {
	owner, err := otc.ownable.Owner(nil)
	if err != nil {
		return "", err
	}

	return owner.Hex(), nil
}
