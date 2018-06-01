package team

import (
	"strings"
	"time"

	"github.com/dungvan/soccer-social-network-api/model"
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/dungvan/soccer-social-network-api/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	GetAllTeam(search string, page uint) (total uint, teams []model.Team, err error)
	FindTeamByID(teamID uint) (*model.Team, error)
	GetAllTeamsByMasterUserID(masterUserID uint) (total uint, teams []model.Team, err error)
	GetAllTeamsByPlayerUserName(userName string) (total uint, teams []model.Team, err error)
	CreateTeam(team *model.Team, transaction *gorm.DB) error
	CreateTeamPlayers(teamPlayers []model.TeamPlayer, transaction *gorm.DB) error
	GetTeamMaster(teamID uint) (*model.User, error)
	GetTeamPlayers(teamID uint) ([]Player, error)
	UpdateTeam(team *model.Team, transaction *gorm.DB) error
	DeleteTeam(teamID uint, transaction *gorm.DB) error
	DeleteTeamMaster(teamID uint, transaction *gorm.DB) error
	DeleteTeamPlayers(teamsID []uint, transaction *gorm.DB) error
	DeleteAllTeamPlayers(teamID uint, transaction *gorm.DB) error
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) GetAllTeam(search string, page uint) (uint, []model.Team, error) {
	var total uint
	teams := make([]model.Team, 0)
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description, teams.created_at").
		Where("teams.name LIKE ?", search+"%%").
		Count(&total).
		Offset(pagingLimit * (page - 1)).
		Limit(pagingLimit).Order("id asc").
		Scan(&teams).Error
	return total, teams, utils.ErrorsWrap(err, "can't get all teams")
}

func (r *repository) FindTeamByID(teamID uint) (*model.Team, error) {
	team := &model.Team{}
	err := r.db.Model(&model.Team{}).Where("id = ?", teamID).First(team).Error
	if err == gorm.ErrRecordNotFound {
		return team, err
	}
	return team, utils.ErrorsWrap(err, "can't find team")
}

func (r *repository) GetAllTeamsByMasterUserID(masterUserID uint) (uint, []model.Team, error) {
	teams := make([]model.Team, 0)
	var total uint
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description, teams.created_at").
		Joins(`INNER JOIN masters ON (masters.owner_type = 'teams' AND masters.owner_id = teams.id AND masters.user_id = ? AND masters.deleted_at IS NULL)`, masterUserID).
		Count(&total).
		Order("teams.created_at desc, teams.id desc").
		Scan(&teams).Error
	return total, teams, utils.ErrorsWrap(err, "can't get team.")
}

func (r *repository) GetAllTeamsByPlayerUserName(userName string) (uint, []model.Team, error) {
	teams := make([]model.Team, 0)
	var total uint
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description, teams.created_at").
		Joins(`INNER JOIN team_players ON (team_players.team_id = teams.id AND team_players.deleted_at IS NULL)`).
		Joins(`INNER JOIN users ON (users.deleted_at IS NULL AND team_players.user_id = users.id AND users.user_name = ?)`, userName).
		Count(&total).
		Order("teams.created_at desc, teams.id desc").
		Scan(&teams).Error
	return total, teams, utils.ErrorsWrap(err, "can't get team.")
}

func (r *repository) CreateTeam(t *model.Team, tx *gorm.DB) error {
	result := tx.Create(t)
	return utils.ErrorsWrap(result.Error, "can't create team")
}

func (r *repository) CreateTeamPlayers(teamPlayers []model.TeamPlayer, tx *gorm.DB) error {
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
		Select("users.id, users.user_name, users.first_name, users.last_name").
		Joins(`INNER JOIN masters ON (masters.user_id = users.id  AND masters.deleted_at IS NULL)`).
		Joins(`INNER JOIN teams ON (teams.id = masters.owner_id AND masters.owner_type = 'teams' AND teams.id = ? AND teams.deleted_at IS NULL)`, teamID).
		Scan(&mstUser)
	return &mstUser, utils.ErrorsWrap(result.Error, "can't get team-master relation")
}

func (r *repository) GetTeamPlayers(teamID uint) ([]Player, error) {
	players := make([]Player, 0)
	result := r.db.Model(&model.User{}).
		Select("users.id, users.first_name, users.last_name, users.user_name, users.score, team_players.position").
		Joins(`INNER JOIN team_players ON team_players.user_id = users.id AND team_players.team_id = ? AND team_players.deleted_at IS NULL`, teamID).
		Scan(&players)
	return players, utils.ErrorsWrap(result.Error, "can't get team-players relation")
}

func (r *repository) UpdateTeam(team *model.Team, transaction *gorm.DB) error {
	return utils.ErrorsWrap(transaction.Model(&model.Team{}).Update(team).Error, "can't update team")
}

func (r *repository) DeleteTeam(teamID uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(transaction.Where("id = ?", teamID).Delete(&model.Team{}).Error, "can't delete team")
}

func (r *repository) DeleteTeamMaster(teamID uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(transaction.Where("owner_type = 'teams' AND owner_id = ?", teamID).Delete(&model.Master{}).Error, "can't delete related team-master")
}

func (r *repository) DeleteTeamPlayers(teamsID []uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(transaction.Where("team_id IN (?)", teamsID).Delete(&model.TeamPlayer{}).Error, "can't delete related team-players")
}

func (r *repository) DeleteAllTeamPlayers(teamID uint, transaction *gorm.DB) error {
	return utils.ErrorsWrap(transaction.Where("team_id = ?", teamID).Delete(&model.TeamPlayer{}).Error, "can't delete all related team-players")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
