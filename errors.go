package gormrepo

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ErrorOrNotExist wraps an error with additional flag indicating record not found
// Provides convenient way to distinguish between actual errors and missing records
// Used by FirstE method to return structured error information
//
// ErrorOrNotExist 封装错误并附加标志指示记录是否未找到
// 提供便捷方式区分实际错误和记录缺失
// 被 FirstE 方法用于返回结构化的错误信息
type ErrorOrNotExist struct {
	Cause    error // Original error // 原始错误
	NotExist bool  // True when record not found // 记录未找到时为 true
}

// NewErrorOrNotExist creates ErrorOrNotExist from the given error
// Checks if error is gorm.ErrRecordNotFound and sets NotExist flag
//
// NewErrorOrNotExist 从给定错误创建 ErrorOrNotExist
// 检查错误是否为 gorm.ErrRecordNotFound 并设置 NotExist 标志
func NewErrorOrNotExist(cause error) *ErrorOrNotExist {
	return &ErrorOrNotExist{
		Cause:    cause,
		NotExist: errors.Is(cause, gorm.ErrRecordNotFound),
	}
}
