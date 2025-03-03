package entities

import "github.com/pkg/errors"

// invalid parameters, internal и т.п.

var (
	ErrInvalidParameter = errors.New("invalid parameter")
	ErrStorageGetFailed = errors.New("failed to get data from storage")
)
