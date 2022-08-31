package http

import "errors"

var (
	BadRequestError        = errors.New("bad request")
	NotFoundError          = errors.New("not found")
	RateLimitError         = errors.New("rate limit exceeded")
	InternalServerError    = errors.New("internal server error")
	ServerDownError        = errors.New("server is down")
	UnknownStatusCodeError = errors.New("unknown status code")
)
