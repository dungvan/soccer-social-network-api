package user

import (
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	Register() error
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}
