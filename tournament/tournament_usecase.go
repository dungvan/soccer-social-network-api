package tournament

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Create a tournament
	Create(CreateRequest) (tournamentID uint, err error)
	// Show a tournament
	Show(tournamentID uint) (RespTournament, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Create(r CreateRequest) (tournamentID uint, err error) {
	isError := false
	tournament := &model.Tournament{
		Master:      &model.Master{UserID: r.UserID},
		Description: r.Description,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
	}
	tx := u.db.Begin()
	defer func() {
		if isError {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := u.repository.CreateTournament(tournament, tx); err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreateTournament() error")
	}
	if err := u.repository.CreateTournamentTeams(tx, tournament.ID, r.Teams...); err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreateTournamentTeam() error")
	}
	return tournament.ID, nil
}

func (u *usecase) Show(tournamentID uint) (RespTournament, error) {
	tournament, err := u.repository.FindTournamentByID(tournamentID)
	if err == gorm.ErrRecordNotFound {
		return RespTournament{TypeOfStatusCode: http.StatusNotFound}, utils.ErrorsNew("the Tournament dose not exist")
	} else if err != nil {
		return RespTournament{}, utils.ErrorsWrap(err, "repository.FindtournamentByID error")
	}
	mstUser, err := u.repository.GetTournamentMaster(tournament.ID)
	if err != nil {
		return RespTournament{}, utils.ErrorsWrap(err, "repositiory.GetTournamentMaster() error")
	}
	teams, err := u.repository.GetTournamentTeams(tournament.ID)
	if err != nil {
		return RespTournament{}, utils.ErrorsWrap(err, "repositiory.GetTournamentTeams() error")
	}
	respTournamentData := RespTournament{
		ID: tournament.ID,
		Master: RespMaster{
			ID:       mstUser.ID,
			UserName: mstUser.UserName,
			FullName: mstUser.FullName,
		},
		Description: tournament.Description,
		StartDate:   tournament.StartDate,
		EndDate:     tournament.EndDate,
		Teams:       teams,
	}

	return respTournamentData, nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
