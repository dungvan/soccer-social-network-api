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
	GetTeamByID(id uint) (RespTeam, error)
	FindMatchByID(matchID uint) (*model.Match, error)
	GetTeamsIDByPlayerUserName(userName string) ([]uint, error)
	GetMatchesByTeamsID(teamsID []uint) (total uint, matches []model.Match, err error)
	GetMatchesByMaster(masterUserID uint) ([]model.Match, error)
	GetMatchMaster(matchID uint) (*model.User, error)
	CreateMatch(match *model.Match, transaction *gorm.DB) error
	GetAllMatches(page uint) (total uint, matches []model.Match, err error)
}

type repository struct {
	base.Repository
	db    *gorm.DB
	redis *redis.Conn
}

func (r *repository) GetAllMatches(page uint) (total uint, matches []model.Match, err error) {
	matches = make([]model.Match, 0)
	total = 0
	result := r.db.Model(&model.Match{}).Count(&total)
	err = result.Offset(pagingLimit * (page - 1)).
		Limit(pagingLimit).Order("created_at desc, id desc").
		Scan(&matches).Error
	return total, matches, utils.ErrorsWrap(err, "can't get all matches")
}

func (r *repository) FindMatchByID(matchID uint) (*model.Match, error) {
	match := &model.Match{}
	err := r.db.Model(&model.Match{}).Where("id = ?", matchID).First(match).Error
	if err == gorm.ErrRecordNotFound {
		return match, err
	}
	return match, utils.ErrorsWrap(err, "can't find match")
}

func (r *repository) GetTeamByID(id uint) (RespTeam, error) {
	team := RespTeam{}
	err := r.db.Model(&model.Team{}).Select(`id, name, description`).Where(`id = ?`, id).Scan(&team).Error
	if err != nil {
		return team, utils.ErrorsWrap(err, "can't get Team by id")
	}
	team.Players, err = r.GetTeamPlayers(team.ID)
	if err != nil {
		return team, utils.ErrorsWrap(err, "GetTeamPlayers() error")
	}
	master, err := r.GetTeamMaster(team.ID)
	team.Master = RespMaster{
		ID:        master.ID,
		UserName:  master.UserName,
		FirstName: master.FirstName,
		LastName:  master.LastName,
	}
	return team, utils.ErrorsWrap(err, "GetTeamMaster() error")
}

func (r *repository) GetTeamsIDByPlayerUserName(userName string) ([]uint, error) {
	teamsID := make([]uint, 0)
	rows, err := r.db.Model(&model.TeamPlayer{}).
		Select("team_players.team_id").
		Joins(`INNER JOIN users ON (users.id = team_players.user_id AND users.deleted_at IS NULL AND users.user_name = ?)`, userName).
		Rows()
	defer rows.Close()
	for rows.Next() {
		var id uint
		rows.Scan(&id)
		teamsID = append(teamsID, id)
	}
	return teamsID, utils.ErrorsWrap(err, "can't get teams id.")
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

func (r *repository) GetMatchesByTeamsID(teamsID []uint) (uint, []model.Match, error) {
	matches := make([]model.Match, 0)
	var total uint
	err := r.db.Model(&model.Match{}).
		Where(`team1_id IN (?) OR team2_id IN (?)`, teamsID, teamsID).
		Count(&total).
		Scan(&matches).Error
	return total, matches, utils.ErrorsWrap(err, "can't get matches by array of teams id.")
}

func (r *repository) GetMatchesByMaster(masterUserID uint) ([]model.Match, error) {
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
		Select("users.id, users.user_name, users.first_name, users.last_name").
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
