package match

import (
	"github.com/dungvan2512/soccer-social-network-api/model"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	"github.com/dungvan2512/soccer-social-network-api/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	FindMatchByID(matchID uint) (*model.Match, error)
	GetAllMatchesByMasterUserID(masterUserID uint) ([]model.Match, error)
	GetMatchMaster(matchID uint) (*model.User, error)
	CreateMatch(match *model.Match, transaction *gorm.DB) error
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) FindMatchByID(matchID uint) (*model.Match, error) {
	match := &model.Match{}
	err := r.db.Model(&model.Match{}).Where("id = ?", matchID).First(match).Error
	if err == gorm.ErrRecordNotFound {
		return match, err
	}
	return match, utils.ErrorsWrap(err, "can't find match")
}

func (r *repository) GetAllMatchesByMasterUserID(masterUserID uint) ([]model.Match, error) {
	matches := make([]model.Match, 0)
	err := r.db.Model(&model.Match{}).
		Select("matches.id, matches.description, matches.start_date").
		Joins(`INNER JOIN masters ON (masters.owner_type = 'matches' AND masters.owner_id = matches.id AND masters.user_id = ? AND masters.deleted_at IS NULL)`, masterUserID).
		Limit(100).
		Order("matches.start_date desc, matches.id desc").
		Scan(&matches).Error
	return matches, utils.ErrorsWrap(err, "can't get match.")
}

func (r *repository) GetMatchMaster(matchID uint) (*model.User, error) {
	mstUser := model.User{}
	result := r.db.Model(&mstUser).
		Select("users.id, users.user_name, users.full_name").
		Joins(`INNER JOIN masters ON (masters.user_id = users.id AND masters.deleted_at IS NULL)`).
		Joins(`INNER JOIN matches ON (matches.id = masters.owner_id AND masters.owner_type = 'matches' AND matches.id = ? AND matches.deleted_at IS NULL)`, matchID).
		Scan(&mstUser)
	return &mstUser, utils.ErrorsWrap(result.Error, "can't get match-master relation")
}

func (r *repository) CreateMatch(m *model.Match, tx *gorm.DB) error {
	result := tx.Create(m)
	return utils.ErrorsWrap(result.Error, "can't create match")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
