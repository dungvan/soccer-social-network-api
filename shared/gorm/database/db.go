package database

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// QueryErrorCallback returns error when select master database error.
func QueryErrorCallback(scope *gorm.Scope) {
	_ = scope.Err(errors.New("can't query master database"))
}

// RowQueryErrorCallback returns error when select master database error.
func RowQueryErrorCallback(scope *gorm.Scope) {
	_ = scope.Err(errors.New("can't row_query master database"))
}

// CreateErrorCallback returns error when create read replica database error.
func CreateErrorCallback(scope *gorm.Scope) {
	_ = scope.Err(errors.New("can't create read replica database"))
}

// UpdateErrorCallback returns error when update read replica database error.
func UpdateErrorCallback(scope *gorm.Scope) {
	_ = scope.Err(errors.New("can't update read replica database"))
}

// DeleteErrorCallback returns error when delete read replica database error.
func DeleteErrorCallback(scope *gorm.Scope) {
	_ = scope.Err(errors.New("can't delete read replica database"))
}
