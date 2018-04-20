package match

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Create a match
	Create(CreateRequest) (matchID uint, err error)
	// Show a match
	Show(matchID uint) (RespMatch, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Create(r CreateRequest) (matchID uint, err error) {
	isError := false
	match := &model.Match{
		Master:      &model.Master{UserID: r.UserID},
		Description: r.Description,
		StartDate:   r.StartDate,
		Team1ID:     r.Team1ID,
		Team2ID:     r.Team2ID,
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
		return RespMatch{TypeOfStatusCode: http.StatusNotFound}, utils.ErrorsNew("the Match dose not exist")
	} else if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repository.FindmatchByID error")
	}
	mstUser, err := u.repository.GetMatchMaster(match.ID)
	if err != nil {
		return RespMatch{}, utils.ErrorsWrap(err, "repositiory.GetMatchMaster() error")
	}
	respMatchData := RespMatch{
		ID: match.ID,
		Master: RespMaster{
			ID:       mstUser.ID,
			UserName: mstUser.UserName,
			FullName: mstUser.FullName,
		},
		Description: match.Description,
		StartDate:   match.StartDate,
		Team1ID:     match.Team1ID,
		Team2ID:     match.Team2ID,
	}

	return respMatchData, nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
