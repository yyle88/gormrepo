package gormrepo

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewErrorOrNotExist(t *testing.T) {
	{
		erb := NewErrorOrNotExist(errors.New("wrong"))
		require.NotErrorIs(t, erb.Cause, gorm.ErrRecordNotFound)
		require.False(t, erb.NotExist)
	}
	{
		erb := NewErrorOrNotExist(gorm.ErrRecordNotFound)
		require.ErrorIs(t, erb.Cause, gorm.ErrRecordNotFound)
		require.True(t, erb.NotExist)
	}
}
