package error

import "errors"

var (
	ServerErr        = errors.New("server error")
	EmptyBodyRequest = errors.New("empty body request")
)
