package user

import (
	"net/http"
	"strings"

	"github.com/dungvan/soccer-social-network-api/model"
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/dungvan/soccer-social-network-api/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Register usecase
	Register(RegisterReuqest) error
	// Login usecase
	Login(LoginRequest) (LoginResponse, error)
	// SendFriendRequest usecase
	SendFriendRequest(FriendRequest) error
	// Show a user
	Show(userName string) (RespUser, error)
	// Index get all user  by super admin
	Index(page uint) (IndexResponse, error)
	// Delete a user
	Delete(id uint) error
	// Update a user
	Update(UpdateRequest) (RespUser, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Register(r RegisterReuqest) error {
	user := model.User{UserName: r.UserName, Email: r.Email, Password: r.Password, FirstName: r.FirstName, LastName: r.LastName}.HashAndSaltPassword()
	err := u.repository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecase) Login(l LoginRequest) (LoginResponse, error) {
	var user *model.User
	var err error
	var token string
	if strings.ContainsAny(l.UserNameOrEmail, "@ & .") {
		user, err = u.repository.FindUserByEmail(l.UserNameOrEmail)
	} else {
		user, err = u.repository.FindUserByUserName(l.UserNameOrEmail)
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return LoginResponse{}, errUserNameOrPassword
		}
		return LoginResponse{}, err
	}
	if ok := u.repository.CheckLogin(*user, l.Password); ok {
		// store user to JWT
		token, err = u.repository.GenerateToken(user)
		if err != nil {
			return LoginResponse{}, utils.ErrorsWrap(err, "repository.GenerateToken() error")
		}
	} else {
		return LoginResponse{}, errUserNameOrPassword
	}
	return LoginResponse{ID: user.ID, UserName: user.UserName, Role: user.Role, Token: token}, nil
}

func (u *usecase) SendFriendRequest(FriendRequest) error {
	return nil
}

func (u *usecase) Show(userName string) (RespUser, error) {
	user, err := u.repository.FindUserByUserName(userName)
	if err == gorm.ErrRecordNotFound {
		return RespUser{TypeOfStatusCode: http.StatusNotFound}, utils.ErrorsNew("the User dose not exist")
	} else if err != nil {
		return RespUser{}, utils.ErrorsWrap(err, "repository.FinduserByID error")
	}
	respUserData := RespUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Birthdate: user.Birthday,
		City:      user.City,
		Country:   user.Country,
		About:     user.About,
		Quote:     user.Quote,
		Role:      user.Role,
	}

	return respUserData, nil
}

func (u *usecase) Index(page uint) (IndexResponse, error) {
	if page < 1 {
		page = 1
	}
	total, users, err := u.repository.GetAllUser(page)
	if err == gorm.ErrRecordNotFound {
		return IndexResponse{Users: []RespUser{}}, nil
	} else if err != nil {
		return IndexResponse{Total: total, Users: []RespUser{}}, utils.ErrorsWrap(err, "repository.GetAllUser() error")
	}
	response := IndexResponse{
		Total: total,
		Users: func() []RespUser {
			respUsers := make([]RespUser, 0)
			for _, user := range users {
				respUsers = append(respUsers, RespUser{
					ID:        user.ID,
					UserName:  user.UserName,
					Email:     user.Email,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Birthdate: user.Birthday,
					City:      user.City,
					Country:   user.Country,
					About:     user.About,
					Quote:     user.Quote,
					Role:      user.Role,
				})
			}
			return respUsers
		}(),
	}
	return response, err
}

func (u *usecase) Update(r UpdateRequest) (RespUser, error) {
	user, err := u.repository.FindUserByID(r.ID)
	if err != nil {
		return RespUser{}, utils.ErrorsWrap(err, "repository.GetAllUser() error")
	}
	user.FirstName = r.FirstName
	user.LastName = r.LastName
	user.City = r.City
	user.Country = r.Country
	user.About = r.About
	user.Birthday = r.Birthday
	user.Quote = r.Quote
	if r.Password != "" {
		user.Password = r.Password
		*user = user.HashAndSaltPassword()
	}
	err = u.repository.UpdateUser(user)
	if err != nil {
		return RespUser{}, utils.ErrorsWrap(err, "repository.UpdateUser() error")
	}
	userResp := RespUser{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Birthdate: user.Birthday,
		City:      user.City,
		Country:   user.Country,
		About:     user.About,
		Quote:     user.Quote,
		Role:      user.Role,
	}
	return userResp, err
}

func (u *usecase) Delete(id uint) error {
	return u.repository.DeleteUser(id)
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *base.Usecase, master *gorm.DB, r Repository) Usecase {
	return &usecase{Usecase: *bu, db: master, repository: r}
}
