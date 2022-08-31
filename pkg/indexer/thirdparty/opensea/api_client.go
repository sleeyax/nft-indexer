package opensea

import (
	"encoding/json"
	"fmt"
	"github.com/sleeyax/gotcha"
	"nft-indexer/pkg/indexer/http"
)

type ApiClient struct {
	client *gotcha.Client
}

func NewOpenSea(apiKey string) (*ApiClient, error) {
	client, err := http.NewDefaultHttpClient("OpenSea", "https://api.opensea.io/api/v1/", apiKey)
	return &ApiClient{client}, err
}

// GetNFTCollection reads an NFT collection (including its metadata) from OpenSea.
func (osp *ApiClient) GetNFTCollection(address string) (*ContractResponse, error) {
	res, err := osp.client.Get(fmt.Sprintf("asset_contract/%s", address))
	if err != nil {
		return nil, err
	}

	jsonString, err := res.Text()
	if err != nil {
		return nil, err
	}

	var contractResponse ContractResponse
	if err = json.Unmarshal([]byte(jsonString), &contractResponse); err != nil {
		return nil, err
	}

	return &contractResponse, nil
}

// GetTotalSupply reads the total amount of tokens in supply.
func (osp *ApiClient) GetTotalSupply(openSeaSlug string) int {
	return 0
}
