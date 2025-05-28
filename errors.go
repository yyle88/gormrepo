package gormrepo

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ErrorOrNotExist struct {
	Cause    error
	NotExist bool
}

func NewErrorOrNotExist(cause error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Cause:    cause,
		NotExist: errors.Is(cause, gorm.ErrRecordNotFound),
	}
}
