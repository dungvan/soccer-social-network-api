package user

import (
	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	Register(RegisterReuqest) error
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Register(r RegisterReuqest) error {
	user := model.User{UserName: r.UserName, Email: r.Email, Password: r.Password, FullName: r.FullName, Birthday: r.Birthday}.HashAndSaltPassword()
	err := u.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *base.Usecase, master *gorm.DB, r Repository) Usecase {
	return &usecase{Usecase: *bu, db: master, repository: r}
}
