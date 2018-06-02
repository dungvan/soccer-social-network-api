package tournament

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
	FindTournamentByID(tournamentID uint) (*model.Tournament, error)
	GetAllTournaments(search string, page uint) (total uint, tournaments []model.Tournament, err error)
	GetAllTournamentsByMaster(masterUserID uint) (total uint, tournaments []model.Tournament, err error)
	GetTournamentMaster(tournamentID uint) (*model.User, error)
	CreateTournament(tournament *model.Tournament, transaction *gorm.DB) error
	CreateTournamentTeams(transaction *gorm.DB, tournamentID uint, teams ...ReqTeam) error
	GetTournamentTeams(tournamentID uint) ([]Team, error)
	GetTournamentMatches(tournamentID uint) ([]Match, error)
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) GetAllTournaments(search string, page uint) (total uint, tournaments []model.Tournament, err error) {
	tournaments = make([]model.Tournament, 0)
	err = r.db.Model(&model.Tournament{}).Where(`name LIKE ?`, search+"%%").
		Count(&total).
		Offset(pagingLimit * (page - 1)).
		Limit(pagingLimit).
		Scan(&tournaments).Error
	return total, tournaments, utils.ErrorsWrap(err, "can't get tournaments")
}

func (r *repository) FindTournamentByID(tournamentID uint) (*model.Tournament, error) {
	tournament := &model.Tournament{}
	err := r.db.Model(&model.Tournament{}).Where("id = ?", tournamentID).First(tournament).Error
	if err == gorm.ErrRecordNotFound {
		return tournament, err
	}
	return tournament, utils.ErrorsWrap(err, "can't find tournament")
}

func (r *repository) GetAllTournamentsByMaster(masterUserID uint) (total uint, tournaments []model.Tournament, err error) {
	tournaments = make([]model.Tournament, 0)
	err = r.db.Model(&model.Tournament{}).
		Select("tournaments.id, tournaments.name, tournaments.description, tournaments.start_date, tournaments.end_date").
		Joins(`INNER JOIN masters ON (masters.owner_type = 'tournaments' AND masters.owner_id = tournaments.id AND masters.user_id = ? AND masters.deleted_at IS NULL)`, masterUserID).
		Count(&total).
		Order("tournaments.start_date desc, tournaments.id desc").
		Scan(&tournaments).Error
	return total, tournaments, utils.ErrorsWrap(err, "can't get tournament.")
}

func (r *repository) GetTournamentMaster(tournamentID uint) (*model.User, error) {
	mstUser := model.User{}
	result := r.db.Model(&mstUser).
		Select("users.id, users.user_name, users.first_name, users.last_name").
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
	err := r.db.Model(&model.Team{}).
		Select("teams.id, teams.name, teams.description").
		Joins(`INNER JOIN tournament_teams ON tournament_teams.team_id = teams.id AND tournament_teams.tournament_id = ? AND tournament_teams.deleted_at IS NULL`, tournamentID).
		Scan(&teams).Error
	if err != nil {
		return teams, utils.ErrorsWrap(err, "can't get team-tournament_teams relation")
	}
	for index, team := range teams {
		err = r.db.Model(&model.User{}).Select("DISTINCT users.id, users.user_name, users.first_name, users.last_name").
			Joins(`INNER JOIN masters ON (masters.user_id = users.id AND masters.deleted_at IS NULL)`).
			Joins(`INNER JOIN team_players ON (masters.owner_type = 'teams' AND masters.owner_id = ? AND team_players.deleted_at IS NULL)`, team.ID).
			Scan(&teams[index].Master).Error
		if err != nil {
			return teams, utils.ErrorsWrap(err, "can't get team masters relation")
		}
		teams[index].Players = make([]Player, 0)
		err = r.db.Model(&model.User{}).Select("DISTINCT users.id, users.user_name, users.first_name, users.last_name, team_players.position").
			Joins(`INNER JOIN team_players ON (users.id = team_players.user_id AND team_players.team_id = ? AND team_players.deleted_at IS NULL)`, team.ID).
			Scan(&teams[index].Players).Error
		if err != nil {
			return teams, utils.ErrorsWrap(err, "can't get team players relation")
		}
	}
	return teams, nil
}

func (r *repository) GetTournamentMatches(tournamentID uint) ([]Match, error) {
	matches := make([]Match, 0)
	err := r.db.Model(&model.Match{}).Select(`id, description, start_date, team1_id, team2_id, team1_goals, team2_goals`).
		Where(`tournament_id = ?`, tournamentID).
		Scan(&matches).Error
	return matches, utils.ErrorsWrap(err, "can't get tournament matches")
}

// NewRepository create new instance of Repository
func NewRepository(br *base.Repository, db *gorm.DB, redis *redis.Conn) Repository {
	return &repository{*br, db, redis}
}
