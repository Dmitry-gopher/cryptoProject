package entities

import "github.com/pkg/errors"

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrStorageGetFailed = errors.New("failed to get data from storage")
	ErrStorage          = errors.New("something went wrong")
	ErrInternal         = errors.New("some internal error")
)
