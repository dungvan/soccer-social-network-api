package team

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Index usecase
	Index(userID uint) (IndexResponse, error)
	// Create a team
	Create(CreateRequest) (teamID uint, err error)
	// Show a team
	Show(teamID uint) (RespTeam, error)
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Index(userID uint) (IndexResponse, error) {
	response := IndexResponse{
		Master: RespTeams{
			Teams: []RespTeam{},
		},
		Player: RespTeams{
			Teams: []RespTeam{},
		},
	}
	master, err := u.repository.GetAllTeamsByMasterUserID(userID)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GetAllTeamsByMasterUserID() error")
	}
	for _, team := range master {
		mstUser, err := u.repository.GetTeamMaster(team.ID)
		if err != nil {
			return response, utils.ErrorsWrap(err, "repositiory.GetTeamMaster() error")
		}
		pls, err := u.repository.GetTeamPlayers(team.ID)
		if err != nil {
			return response, utils.ErrorsWrap(err, "repositiory.GetRelatedTeamPlayers() error")
		}
		respTeamData := RespTeam{
			ID: team.ID,
			Master: RespMaster{
				ID:       mstUser.ID,
				UserName: mstUser.UserName,
				FullName: mstUser.FullName,
			},
			Description: team.Description,
			Name:        team.Name,
			Players:     pls,
		}
		response.Master.Teams = append(response.Master.Teams, respTeamData)
	}
	players, err := u.repository.GetAllTeamsByPlayerUserID(userID)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GetAllTeamsByPlayerUserID() error")
	}
	for _, team := range players {
		mstUser, err := u.repository.GetTeamMaster(team.ID)
		if err != nil {
			return response, utils.ErrorsWrap(err, "repositiory.GetTeamMaster() error")
		}
		pls, err := u.repository.GetTeamPlayers(team.ID)
		if err != nil {
			return response, utils.ErrorsWrap(err, "repositiory.GetRelatedTeamPlayers() error")
		}
		respTeamData := RespTeam{
			ID: team.ID,
			Master: RespMaster{
				ID:       mstUser.ID,
				UserName: mstUser.UserName,
				FullName: mstUser.FullName,
			},
			Description: team.Description,
			Name:        team.Name,
			Players:     pls,
		}
		response.Player.Teams = append(response.Player.Teams, respTeamData)
	}
	response.Master.ResultCount = len(master)
	response.Player.ResultCount = len(players)
	return response, nil
}

func (u *usecase) Create(r CreateRequest) (uint, error) {
	isError := false
	team := &model.Team{}
	team.Name = r.Name
	team.Master = &model.Master{UserID: r.UserID}
	team.Description = r.Description
	tx := u.db.Begin()
	defer func() {
		if isError {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := u.repository.CreateTeam(team, tx); err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreateTeam() error")
	}

	teamplayers := make([]model.TeamPlayer, 0)

	for _, player := range r.Players {
		teamplayers = append(teamplayers, model.TeamPlayer{TeamID: team.ID, UserID: player.ID, Position: player.Position})
	}

	if err := u.repository.CreateTeamPlayer(teamplayers, tx); err != nil {
		isError = true
		return 0, utils.ErrorsWrap(err, "repository.CreateTeamPlayer() error")
	}

	return team.ID, nil
}

func (u *usecase) Show(teamID uint) (RespTeam, error) {
	team, err := u.repository.FindTeamByID(teamID)
	if err == gorm.ErrRecordNotFound {
		return RespTeam{TypeOfStatusCode: http.StatusNotFound}, utils.ErrorsNew("the Team dose not exist")
	} else if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repository.FindteamByID error")
	}
	mstUser, err := u.repository.GetTeamMaster(team.ID)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repositiory.GetTeamMaster() error")
	}
	pls, err := u.repository.GetTeamPlayers(team.ID)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repositiory.GetRelatedTeamPlayers() error")
	}
	respTeamData := RespTeam{
		ID: team.ID,
		Master: RespMaster{
			ID:       mstUser.ID,
			UserName: mstUser.UserName,
			FullName: mstUser.FullName,
		},
		Description: team.Description,
		Name:        team.Name,
		Players:     pls,
	}

	return respTeamData, nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
