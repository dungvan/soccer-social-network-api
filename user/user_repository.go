package user

import (
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	SampleRepository()
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}
