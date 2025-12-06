package gormrepo

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// TestNewErrorOrNotExist tests ErrorOrNotExist creation and ErrRecordNotFound detection
// TestNewErrorOrNotExist 测试 ErrorOrNotExist 创建和 ErrRecordNotFound 检测
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
