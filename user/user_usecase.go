package user

import (
	"strings"

	"github.com/dungvan2512/socker-social-network/model"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Register usecase
	Register(RegisterReuqest) error
	// Login usecase
	Login(LoginRequest) (token string, err error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Register(r RegisterReuqest) error {
	user := model.User{UserName: r.UserName, Email: r.Email, Password: r.Password, FullName: r.FullName, Birthday: r.Birthday}.HashAndSaltPassword()
	err := u.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) Login(l LoginRequest) (string, error) {
	var user *model.User
	var err error
	var token string
	if strings.ContainsAny(l.UserNameOrEmail, "@ & .") {
		user, err = u.repository.FindUserByEmail(l.UserNameOrEmail)
	} else {
		user, err = u.repository.FindUserByUserName(l.UserNameOrEmail)
	}
	if err == gorm.ErrRecordNotFound {
		return "", errUserNameOrPassword
	}
	if ok := u.repository.CheckLogin(*user, l.Password); ok {
		// store user to JWT
		token, err = u.repository.GenerateToken(user)
		if err != nil {
			return "", utils.ErrorsWrap(err, "repository.GenerateToken() error")
		}
	} else {
		return "", errUserNameOrPassword
	}
	return token, nil
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *base.Usecase, master *gorm.DB, r Repository) Usecase {
	return &usecase{Usecase: *bu, db: master, repository: r}
}
