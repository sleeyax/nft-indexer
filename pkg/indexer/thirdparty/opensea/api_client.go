package opensea

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sleeyax/gotcha"
	"net/http"
	"time"
)

type ApiClient struct {
	client *gotcha.Client
}

const (
	openSeaApiUrl = "https://api.opensea.io/api/v1/"
	errorFormat   = "OpenSea provider: [%w]"
)

var (
	BadRequestError        = errors.New("bad request")
	NotFoundError          = errors.New("not found")
	RateLimitError         = errors.New("rate limit exceeded")
	InternalServerError    = errors.New("internal server error")
	ServerDownError        = errors.New("server is down")
	UnknownStatusCodeError = errors.New("unknown status code")
)

func NewOpenSea(apiKey string) (*ApiClient, error) {
	client, err := gotcha.NewClient(&gotcha.Options{
		PrefixURL: openSeaApiUrl,
		Headers: http.Header{
			"x-api-key": {apiKey},
		},
		Timeout:        time.Second * 20,
		FollowRedirect: false,
		Retry:          true,
		RetryOptions: &gotcha.RetryOptions{
			Limit: 3,
		},
		Hooks: gotcha.Hooks{
			BeforeRetry: []gotcha.BeforeRetryHook{
				func(options *gotcha.Options, error error, retryCount int) {
					if error == RateLimitError {
						time.Sleep(time.Second * 1)
					} else if error == ServerDownError {
						time.Sleep(time.Second * 5)
					} else if error == UnknownStatusCodeError {
						time.Sleep(time.Second * 2)
					}
				},
			},
			AfterResponse: []gotcha.AfterResponseHook{
				func(res *gotcha.Response, retry gotcha.RetryFunc) (*gotcha.Response, error) {
					if res.StatusCode == 200 {
						return res, nil
					}

					switch res.StatusCode {
					case 400:
						return nil, fmt.Errorf(errorFormat, BadRequestError)
					case 404:
						return nil, fmt.Errorf(errorFormat, NotFoundError)
					case 429:
						return nil, fmt.Errorf(errorFormat, RateLimitError)
					case 500:
						return nil, fmt.Errorf(errorFormat, InternalServerError)
					case 504:
						return nil, fmt.Errorf(errorFormat, ServerDownError)
					}

					return nil, fmt.Errorf(errorFormat, UnknownStatusCodeError)
				},
			},
		},
	})

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
