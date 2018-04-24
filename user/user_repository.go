package user

import (
	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/auth"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	// Create repo
	CreateUser(model.User) error
	// FindUserByUserName
	FindUserByUserName(userName string) (*model.User, error)
	// FindUserByEmail
	FindUserByEmail(email string) (*model.User, error)
	// CheckLogin return true if password match
	CheckLogin(user model.User, password string) bool
	// GenerateToken for user
	GenerateToken(*model.User) (token string, err error)
	// Create user follow
	CreateUserFollow(userFollow *model.UserFollow) error
	// DeleteUserFollow
	DeleteUserFollow(userFollowID uint) error
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) CreateUser(u model.User) error {
	return r.db.Create(&u).Error
}

func (r *repository) FindUserByUserName(userName string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("user_name = ?", userName).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return user, err
	}
	return user, utils.ErrorsWrap(err, "can't find user")
}

func (r *repository) FindUserByEmail(emmail string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err == gorm.ErrRecordNotFound {
		return user, err
	}
	return user, utils.ErrorsWrap(err, "can't find user")
}

func (r *repository) CheckLogin(user model.User, password string) bool {
	return user.CompareHashAndPassword(password)
}

func (r *repository) GenerateToken(user *model.User) (token string, err error) {
	return auth.GenerateToken(user)
}

func (r *repository) CreateUserFollow(userFollow *model.UserFollow) error {
	return r.db.Create(&userFollow).Error
}

func (r *repository) DeleteUserFollow(userFollowID uint) error {
	return nil
}

// NewRepository responses new Repository instance.
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{Repository: *br, db: db, redis: redis}
}
