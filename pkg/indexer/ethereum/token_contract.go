package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"nft-indexer/pkg/database"
)

var NullAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

type TokenContract struct {
	contract      *Contract
	TokenStandard database.TokenStandard
}

func NewTokenContract(contract *Contract, standard database.TokenStandard) (*TokenContract, error) {
	return &TokenContract{
		contract,
		standard,
	}, nil
}

func (tc *TokenContract) ToOwnable() (*OwnableTokenContract, error) {
	ownable, err := tc.contract.ToOwnable()
	if err != nil {
		return nil, err
	}

	return &OwnableTokenContract{
		TokenContract: tc,
		ownable:       ownable,
	}, nil
}
