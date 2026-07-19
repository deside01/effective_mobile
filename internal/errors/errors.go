package errors

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrFailedQuery  = errors.New("failed to query")
)
