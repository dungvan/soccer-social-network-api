package team

import (
	"strings"
	"time"

	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	FindTeamByID(teamID uint) (*model.Team, error)
	GetAllTeamsByMasterUserID(masterUserID uint) ([]model.Team, error)
	GetAllTeamsByPlayerUserID(playerUserID uint) ([]model.Team, error)
	CreateTeam(team *model.Team, transaction *gorm.DB) error
	CreateTeamPlayer(teamPlayers []model.TeamPlayer, transaction *gorm.DB) error
	GetTeamMaster(teamID uint) (*model.User, error)
	GetTeamPlayers(teamID uint) ([]Player, error)
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) FindTeamByID(teamID uint) (*model.Team, error) {
	team := &model.Team{}
	err := r.db.Model(&model.Team{}).Where("id = ?", teamID).First(team).Error
	if err == gorm.ErrRecordNotFound {
		return team, err
	}
	return team, utils.ErrorsWrap(err, "can't find team")
}

func (r *repository) GetAllTeamsByMasterUserID(masterUserID uint) ([]model.Team, error) {
	teams := make([]model.Team, 0)
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description, teams.created_at").
		Joins(`INNER JOIN masters ON (masters.owner_type = 'teams' AND masters.owner_id = teams.id AND masters.user_id = ? AND masters.deleted_at IS NULL)`, masterUserID).
		Limit(100).
		Order("teams.created_at desc, teams.id desc").
		Scan(&teams).Error
	return teams, utils.ErrorsWrap(err, "can't get team.")
}

func (r *repository) GetAllTeamsByPlayerUserID(playerUserID uint) ([]model.Team, error) {
	teams := make([]model.Team, 0)
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description, teams.created_at").
		Joins(`INNER JOIN team_players ON (team_players.team_id = teams.id AND team_players.deleted_at IS NULL)`).
		Where("team_players.user_id = ?", playerUserID).
		Limit(100).
		Order("teams.created_at desc, teams.id desc").
		Scan(&teams).Error
	return teams, utils.ErrorsWrap(err, "can't get team.")
}

func (r *repository) CreateTeam(t *model.Team, tx *gorm.DB) error {
	result := tx.Create(t)
	return utils.ErrorsWrap(result.Error, "can't create team")
}

func (r *repository) CreateTeamPlayer(teamPlayers []model.TeamPlayer, tx *gorm.DB) error {
	if len(teamPlayers) == 0 {
		return nil
	}
	sqlStr := "INSERT INTO team_players (team_id, user_id, position, created_at, updated_at) VALUES "
	vals := []interface{}{}
	params := []string{}
	for _, teamPlayer := range teamPlayers {
		params = append(params, "(?, ?, ?, ?, ?)")
		vals = append(vals, teamPlayer.TeamID, teamPlayer.UserID, teamPlayer.Position, time.Now(), time.Now())
	}
	sqlStr += strings.Join(params, ",")
	err := tx.Exec(sqlStr, vals...).Error
	return utils.ErrorsWrap(err, "can't create data.")
}

func (r *repository) GetTeamMaster(teamID uint) (*model.User, error) {
	mstUser := model.User{}
	result := r.db.Model(&mstUser).
		Select("users.id, users.user_name, users.full_name").
		Joins(`INNER JOIN masters ON (masters.user_id = users.id  AND masters.deleted_at IS NULL)`).
		Joins(`INNER JOIN teams ON (teams.id = masters.owner_id AND masters.owner_type = 'teams' AND teams.id = ? AND teams.deleted_at IS NULL)`, teamID).
		Scan(&mstUser)
	return &mstUser, utils.ErrorsWrap(result.Error, "can't get team-master relation")
}

func (r *repository) GetTeamPlayers(teamID uint) ([]Player, error) {
	players := make([]Player, 0)
	result := r.db.Model(&model.User{}).
		Select("users.id, users.full_name, users.user_name, users.score, team_players.position").
		Joins(`INNER JOIN team_players ON team_players.user_id = users.id AND team_players.team_id = ? AND team_players.deleted_at IS NULL`, teamID).
		Scan(&players)
	return players, utils.ErrorsWrap(result.Error, "can't get team-players relation")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
