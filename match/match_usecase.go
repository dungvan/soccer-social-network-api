package match

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network-api/model"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	"github.com/dungvan2512/soccer-social-network-api/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Index get all matches
	Index(page uint) (IndexResponse, error)
	// Create a match
	Create(CreateRequest) (matchID uint, err error)
	// Show a match
	Show(matchID uint) (RespMatch, error)
	// GetByUserName return all matches by user name
	GetByUserName(userName string) (IndexResponse, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Index(page uint) (IndexResponse, error) {
	if page < 1 {
		page = 1
	}
	total, matches, err := u.repository.GetAllMatches(page)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetMatchesByTeamsID() error")
	}
	matchesResp := make([]RespMatch, 0)
	for _, match := range matches {
		respMatch := RespMatch{
			ID:          match.ID,
			Description: match.Description,
			StartDate:   match.StartDate,
			Team1Goals:  match.Team1Goals,
			Team2Goals:  match.Team2Goals,
		}
		if match.TournamentID != nil {
			respMatch.Tournament, err = u.repository.GetTournamentByID(*match.TournamentID)
			if err != nil {
				return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTournamentByID() error")
			}
		}
		master, err := u.repository.GetMatchMaster(match.ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetMatchMaster() error")
		}
		respMatch.Master = RespMaster{
			ID:        master.ID,
			UserName:  master.UserName,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}
		respMatch.Team1, err = u.repository.GetTeamByID(match.Team1ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
		}
		respMatch.Team2, err = u.repository.GetTeamByID(match.Team2ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
		}
		matchesResp = append(matchesResp, respMatch)
	}
	return IndexResponse{Total: total, Matches: matchesResp}, nil
}

func (u *usecase) Create(r CreateRequest) (matchID uint, err error) {
	isError := false
	match := &model.Match{
		TournamentID: r.TournamentID,
		Master:       &model.Master{UserID: r.UserID},
		Description:  r.Description,
		StartDate:    r.StartDate,
		Team1ID:      r.Team1ID,
		Team2ID:      r.Team2ID,
	}
	tx := u.db.Begin()
	defer func() {
		if isError {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := u.repository.CreateMatch(match, tx); err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreateMatch() error")
	}
	return match.ID, nil
}

func (u *usecase) Show(matchID uint) (RespMatch, error) {
	match, err := u.repository.FindMatchByID(matchID)
	if err == gorm.ErrRecordNotFound {
		return RespMatch{TypeOfStatusCode: http.StatusNotFound}, utils.ErrorsWrap(nil, "the Match dose not exist")
	} else if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repository.FindmatchByID error")
	}
	mstUser, err := u.repository.GetMatchMaster(match.ID)
	if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repositiory.GetMatchMaster() error")
	}
	team1, err := u.repository.GetTeamByID(match.Team1ID)
	if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
	}
	team2, err := u.repository.GetTeamByID(match.Team2ID)
	if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
	}
	respMatchData := RespMatch{
		ID: match.ID,
		Master: RespMaster{
			ID:        mstUser.ID,
			UserName:  mstUser.UserName,
			FirstName: mstUser.FirstName,
			LastName:  mstUser.LastName,
		},
		Description: match.Description,
		StartDate:   match.StartDate,
		Team1:       team1,
		Team2:       team2,
		Team1Goals:  match.Team1Goals,
		Team2Goals:  match.Team2Goals,
	}
	if match.TournamentID != nil {
		respMatchData.Tournament, err = u.repository.GetTournamentByID(*match.TournamentID)
		if err != nil {
			return RespMatch{}, utils.ErrorsWrap(err, "repository.GetTournamentByID() error")
		}
	}

	return respMatchData, nil
}

func (u *usecase) GetByUserName(userName string) (IndexResponse, error) {
	teamsID, err := u.repository.GetTeamsIDByPlayerUserName(userName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTeamsIDByPlayerUserName() error")
	}
	total, matches, err := u.repository.GetMatchesByTeamsID(teamsID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetMatchesByTeamsID() error")
	}
	matchesResp := make([]RespMatch, 0)
	for _, match := range matches {
		respMatch := RespMatch{
			ID:          match.ID,
			Description: match.Description,
			StartDate:   match.StartDate,
			Team1Goals:  match.Team1Goals,
			Team2Goals:  match.Team2Goals,
		}
		if match.TournamentID != nil {
			respMatch.Tournament, err = u.repository.GetTournamentByID(*match.TournamentID)
			if err != nil {
				return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTournamentByID() error")
			}
		}
		master, err := u.repository.GetMatchMaster(match.ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetMatchMaster() error")
		}
		respMatch.Master = RespMaster{
			ID:        master.ID,
			UserName:  master.UserName,
			FirstName: master.FirstName,
			LastName:  master.LastName,
		}
		respMatch.Team1, err = u.repository.GetTeamByID(match.Team1ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
		}
		respMatch.Team2, err = u.repository.GetTeamByID(match.Team2ID)
		if err != nil {
			return IndexResponse{Matches: []RespMatch{}}, utils.ErrorsWrap(err, "repository.GetTeamByID() error")
		}
		matchesResp = append(matchesResp, respMatch)
	}
	return IndexResponse{Total: total, Matches: matchesResp}, nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
