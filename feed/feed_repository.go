package feed

import (
	"github.com/dungvan2512/soccer-social-network/shared/base"
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

func (r *repository) SampleRepository() {
	return
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
