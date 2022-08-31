package reservoir

import (
	"encoding/json"
	"github.com/Sleeyax/urlValues"
	"github.com/sleeyax/gotcha"
	"nft-indexer/pkg/indexer/http"
	"strconv"
)

type ApiClient struct {
	client *gotcha.Client
}

func NewApiClient(apiKey string) (*ApiClient, error) {
	client, err := http.NewDefaultHttpClient("reservoir", "https://api.reservoir.tools/", apiKey)
	client.Options.Headers.Add("accept", "application/json")
	return &ApiClient{client}, err
}

func (ac *ApiClient) GetTokensInfo(contract string, limit int, cursor string) (*DetailedTokensResponse, error) {
	searchParams := urlValues.Values{
		"contract": {contract},
		"limit":    {strconv.Itoa(limit)},
	}

	if cursor != "" {
		searchParams["continuation"] = append(searchParams["continuation"], cursor)
	}

	res, err := ac.client.Get("tokens/details/v4", &gotcha.Options{
		SearchParams: searchParams,
	})
	if err != nil {
		return nil, err
	}

	jsonString, err := res.Text()
	if err != nil {
		return nil, err
	}

	var tokensResponse DetailedTokensResponse
	if err = json.Unmarshal([]byte(jsonString), &tokensResponse); err != nil {
		return nil, err
	}

	return &tokensResponse, nil
}
