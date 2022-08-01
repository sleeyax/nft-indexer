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

// GetOwner returns the address of the contract owner, which is stored in the contract.
// This value could be different from GetCreator.
func (otc *OwnableTokenContract) GetOwner() (string, error) {
	owner, err := otc.ownable.Owner(nil)
	if err != nil {
		return "", err
	}

	return owner.Hex(), nil
}

// GetCreationEvent returns the transaction (on chain) where the contract was first created.
func (otc *OwnableTokenContract) GetCreationEvent() (*tokens.OwnableOwnershipTransferred, error) {
	iterator, err := otc.ownable.FilterOwnershipTransferred(nil, []common.Address{NullAddress}, []common.Address{})
	if err != nil {
		return nil, err
	}

	if !iterator.Next() {
		return nil, fmt.Errorf("failed to get contract creation event of contract %s on network %s", otc.contract.Address, otc.contract.NetworkId)
	}

	return iterator.Event, nil
}
