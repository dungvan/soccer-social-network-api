package match

import (
	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Create a match
	Create(CreateRequest) (matchID uint, err error)
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

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
