package http

import (
	"fmt"
	"github.com/sleeyax/gotcha"
	"net/http"
	"time"
)

func NewDefaultHttpClient(name string, url string, apiKey string) (*gotcha.Client, error) {
	errorFormat := name + ": [%w]"

	return gotcha.NewClient(&gotcha.Options{
		PrefixURL:      url,
		Timeout:        time.Second * 20,
		FollowRedirect: false,
		Retry:          true,
		Headers: http.Header{
			"x-api-key": {apiKey},
		},
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
}
