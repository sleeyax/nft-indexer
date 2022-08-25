package zora

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sleeyax/gotcha"
	"net/http"
	"text/template"
	"time"
)

const (
	zoraApiUrl  = "https://api.zora.co/graphql"
	errorFormat = "zora provider: [%w]"
)

var (
	TooManyRequestsError   = errors.New("too many requests")
	UnknownStatusCodeError = errors.New("unknown status code")
	InvalidApiKeyError     = errors.New("invalid API key")
)

type GraphQlClient struct {
	client *gotcha.Client
}

func NewGraphQLClient(apiKey string) (*GraphQlClient, error) {
	client, err := gotcha.NewClient(&gotcha.Options{
		PrefixURL: zoraApiUrl,
		Headers: http.Header{
			"X-API-KEY":    {apiKey},
			"Accept":       {"application/json"},
			"Content-Type": {"application/json"},
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
					if error == TooManyRequestsError {
						time.Sleep(time.Second * 1)
					}
				},
			},
			AfterResponse: []gotcha.AfterResponseHook{
				func(res *gotcha.Response, retry gotcha.RetryFunc) (*gotcha.Response, error) {
					if res.StatusCode == 200 {
						return res, nil
					} else if res.StatusCode == 429 {
						return nil, fmt.Errorf(errorFormat, TooManyRequestsError)
					} else if res.StatusCode == 403 {
						return nil, fmt.Errorf(errorFormat, InvalidApiKeyError)
					} else {
						return nil, fmt.Errorf(errorFormat, UnknownStatusCodeError)
					}
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &GraphQlClient{client}, nil
}

func (g *GraphQlClient) parseTemplate(templ string, data map[string]interface{}) (*bytes.Buffer, error) {
	t, err := template.New("").Parse(templ)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	if err = t.Execute(&buffer, data); err != nil {
		return nil, err
	}

	return &buffer, nil
}

func (g *GraphQlClient) AggregateStat(collectionAddress string, topOwnersLimit int) (*GraphQLResponse, error) {
	buffer, err := g.parseTemplate(queryAggregateStatTemplate, map[string]interface{}{
		"CollectionAddress": collectionAddress,
		"TopOwnersLimit":    topOwnersLimit,
	})
	if err != nil {
		return nil, fmt.Errorf(errorFormat, err)
	}

	res, err := g.client.Post("", &gotcha.Options{
		Json: map[string]interface{}{
			"operationName": "MyQuery",
			"query":         buffer.String(),
			"variables":     nil,
		},
	})
	if err != nil {
		return nil, fmt.Errorf(errorFormat, err)
	}

	jsonString, err := res.Text()
	if err != nil {
		return nil, fmt.Errorf(errorFormat, err)
	}

	var graphQLResponse GraphQLResponse
	if err = json.Unmarshal([]byte(jsonString), &graphQLResponse); err != nil {
		return nil, fmt.Errorf(errorFormat, err)
	}

	return &graphQLResponse, nil
}
