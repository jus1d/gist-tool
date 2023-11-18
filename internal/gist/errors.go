package gist

import "errors"

var (
	ErrInvalidRequest = errors.New("can't create a request")
	ErrFileNotExists  = errors.New("file doesn't exists")
	ErrCantReadBody   = errors.New("can't read response body")
)
