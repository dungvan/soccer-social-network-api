package team

import (
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	SampleUsecase()
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) SampleUsecase() {
	return
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
