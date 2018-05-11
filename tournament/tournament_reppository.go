package tournament

import (
	"strings"
	"time"

	"github.com/dungvan2512/soccer-social-network-api/model"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	"github.com/dungvan2512/soccer-social-network-api/shared/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	FindTournamentByID(tournamentID uint) (*model.Tournament, error)
	GetAllTournamentsByMasterUserID(masterUserID uint) ([]model.Tournament, error)
	GetTournamentMaster(tournamentID uint) (*model.User, error)
	CreateTournament(tournament *model.Tournament, transaction *gorm.DB) error
	CreateTournamentTeams(transaction *gorm.DB, tournamentID uint, teams ...ReqTeam) error
	GetTournamentTeams(tournamentID uint) ([]Team, error)
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) FindTournamentByID(tournamentID uint) (*model.Tournament, error) {
	tournament := &model.Tournament{}
	err := r.db.Model(&model.Tournament{}).Where("id = ?", tournamentID).First(tournament).Error
	if err == gorm.ErrRecordNotFound {
		return tournament, err
	}
	return tournament, utils.ErrorsWrap(err, "can't find tournament")
}

func (r *repository) GetAllTournamentsByMasterUserID(masterUserID uint) ([]model.Tournament, error) {
	tournaments := make([]model.Tournament, 0)
	err := r.db.Model(&model.Tournament{}).
		Select("tournaments.id, tournaments.description, tournaments.start_date").
		Joins(`INNER JOIN masters ON (masters.owner_type = 'tournaments' AND masters.owner_id = tournaments.id AND masters.user_id = ? AND masters.deleted_at IS NULL)`, masterUserID).
		Limit(100).
		Order("tournaments.start_date desc, tournaments.id desc").
		Scan(&tournaments).Error
	return tournaments, utils.ErrorsWrap(err, "can't get tournament.")
}

func (r *repository) GetTournamentMaster(tournamentID uint) (*model.User, error) {
	mstUser := model.User{}
	result := r.db.Model(&mstUser).
		Select("users.id, users.user_name, users.full_name").
		Joins(`INNER JOIN masters ON (masters.user_id = users.id AND masters.deleted_at IS NULL)`).
		Joins(`INNER JOIN tournaments ON (tournaments.id = masters.owner_id AND masters.owner_type = 'tournaments' AND tournaments.id = ? AND tournaments.deleted_at IS NULL)`, tournamentID).
		Scan(&mstUser)
	return &mstUser, utils.ErrorsWrap(result.Error, "can't get tournament-master relation")
}

func (r *repository) CreateTournament(m *model.Tournament, tx *gorm.DB) error {
	result := tx.Create(m)
	return utils.ErrorsWrap(result.Error, "can't create tournament")
}

func (r *repository) CreateTournamentTeams(tx *gorm.DB, tournamentID uint, teams ...ReqTeam) error {
	if len(teams) == 0 {
		return nil
	}
	sqlStr := `INSERT INTO tournament_teams(tournament_id, team_id, "group", created_at, updated_at) VALUES `
	vals := []interface{}{}
	params := []string{}
	for _, team := range teams {
		params = append(params, "(?, ?, ?, ?, ?)")
		vals = append(vals, tournamentID, team.ID, team.Group, time.Now(), time.Now())
	}
	sqlStr += strings.Join(params, ",")
	err := tx.Exec(sqlStr, vals...).Error
	return utils.ErrorsWrap(err, "can't create data.")
}

func (r *repository) GetTournamentTeams(tournamentID uint) ([]Team, error) {
	teams := make([]Team, 0)
	result := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, tournament_teams.score, tournament_teams.group").
		Joins(`INNER JOIN tournament_teams ON tournament_teams.team_id = teams.id AND tournament_teams.tournament_id = ? AND tournament_teams.deleted_at IS NULL`, tournamentID).
		Scan(&teams)
	return teams, utils.ErrorsWrap(result.Error, "can't get team-tournament_teams relation")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
