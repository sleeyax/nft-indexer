package ethereum

import "nft-indexer/pkg/indexer"

type Contract struct {
	indexer.Contract
	provider *Provider
}

// TODO: finish & test contract method https://geth.ethereum.org/docs/dapp/native-bindings#accessing-an-ethereum-contract

func (c Contract) NewContract(address string, networkId string, provider *Provider) *Contract {
	return &Contract{
		Contract: indexer.Contract{
			Address:   address,
			NetworkId: networkId,
		},
		provider: provider,
	}
}
