package team

import (
	"net/http"

	"github.com/dungvan/soccer-social-network-api/model"
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/dungvan/soccer-social-network-api/shared/utils"
	"github.com/jinzhu/gorm"
)

// Usecase interface
type Usecase interface {
	// Index usecase
	Index(IndexRequest) (IndexResponse, error)
	// GetByUserName usecase
	GetByUserName(userName string) (IndexResponse, error)
	// GetByMasterID usecase
	GetByMasterID(masterID uint) (IndexResponse, error)
	// Create a team
	Create(CreateRequest) (teamID uint, err error)
	// Show a team
	Show(teamID uint) (RespTeam, error)
	// Update a team
	Update(r UpdateRequest, ctxUser model.User) (RespTeam, error)
	// Delete a team
	Delete(teamID uint, ctxUser model.User) error
}

type usecase struct {
	base.Usecase
	db         *gorm.DB
	repository Repository
}

func (u *usecase) Index(r IndexRequest) (IndexResponse, error) {
	response := IndexResponse{Teams: make([]RespTeam, 0)}
	if r.Page < 1 {
		r.Page = 1
	}
	total, teams, err := u.repository.GetAllTeam(r.Search, r.Page)
	if err == gorm.ErrRecordNotFound {
		return IndexResponse{Teams: []RespTeam{}}, nil
	}
	if err != nil {
		return IndexResponse{Total: total, Teams: []RespTeam{}}, utils.ErrorsWrap(err, "repository.GetAllTeam() error")
	}

	for _, team := range teams {
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
				ID:        mstUser.ID,
				UserName:  mstUser.UserName,
				FirstName: mstUser.FirstName,
				LastName:  mstUser.LastName,
			},
			Description: team.Description,
			Name:        team.Name,
			Players:     pls,
		}
		response.Teams = append(response.Teams, respTeamData)
	}

	response.Total = total
	return response, nil
}

func (u *usecase) GetByMasterID(masterID uint) (IndexResponse, error) {
	response := IndexResponse{Teams: make([]RespTeam, 0)}

	total, teams, err := u.repository.GetAllTeamsByMasterUserID(masterID)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GetAllTeamsByPlayerUserID() error")
	}
	mstUser, err := u.repository.GetTeamMaster(teams[0].ID)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repositiory.GetTeamMaster() error")
	}
	for _, team := range teams {
		pls, err := u.repository.GetTeamPlayers(team.ID)
		if err != nil {
			return response, utils.ErrorsWrap(err, "repositiory.GetRelatedTeamPlayers() error")
		}
		respTeamData := RespTeam{
			ID: team.ID,
			Master: RespMaster{
				ID:        mstUser.ID,
				UserName:  mstUser.UserName,
				FirstName: mstUser.FirstName,
				LastName:  mstUser.LastName,
			},
			Description: team.Description,
			Name:        team.Name,
			Players:     pls,
		}
		response.Teams = append(response.Teams, respTeamData)
	}
	response.Total = total
	return response, nil
}

func (u *usecase) GetByUserName(userName string) (IndexResponse, error) {
	response := IndexResponse{Teams: make([]RespTeam, 0)}

	total, teams, err := u.repository.GetAllTeamsByPlayerUserName(userName)
	if err != nil {
		return response, utils.ErrorsWrap(err, "repository.GetAllTeamsByPlayerUserID() error")
	}
	for _, team := range teams {
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
				ID:        mstUser.ID,
				UserName:  mstUser.UserName,
				FirstName: mstUser.FirstName,
				LastName:  mstUser.LastName,
			},
			Description: team.Description,
			Name:        team.Name,
			Players:     pls,
		}
		response.Teams = append(response.Teams, respTeamData)
	}
	response.Total = total
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

	if err := u.repository.CreateTeamPlayers(teamplayers, tx); err != nil {
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
			ID:        mstUser.ID,
			UserName:  mstUser.UserName,
			FirstName: mstUser.FirstName,
			LastName:  mstUser.LastName,
		},
		Description: team.Description,
		Name:        team.Name,
		Players:     pls,
	}

	return respTeamData, nil
}

func (u *usecase) Update(r UpdateRequest, ctxUser model.User) (RespTeam, error) {
	var err error
	if ctxUser.Role != "s_admin" {
		master, err := u.repository.GetTeamMaster(r.ID)
		if err != nil {
			return RespTeam{}, utils.ErrorsWrap(err, "repository.GetTeamMaster() error")
		}
		if master.ID != ctxUser.ID {
			return RespTeam{}, utils.ErrorsNew("Forbbiden to update the team")
		}
	}

	team, err := u.repository.FindTeamByID(r.ID)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repository.FindTeamByID() error")
	}

	tx := u.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = u.repository.UpdateTeam(team, tx)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repository.UpdateTeam() error")
	}
	err = u.repository.DeleteTeamPlayers(r.PlayersDel, tx)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repository.UpdateTeam() error")
	}

	teamplayers := make([]model.TeamPlayer, 0)
	for _, player := range r.PlayersAdd {
		teamplayers = append(teamplayers, model.TeamPlayer{TeamID: team.ID, UserID: player.ID, Position: player.Position})
	}
	err = u.repository.CreateTeamPlayers(teamplayers, tx)
	if err != nil {
		return RespTeam{}, utils.ErrorsWrap(err, "repository.CreateTeamPlayers() error")
	}

	return u.Show(r.ID)
}

func (u *usecase) Delete(teamID uint, ctxUser model.User) error {
	var err error
	if ctxUser.Role != "s_admin" {
		master, err := u.repository.GetTeamMaster(teamID)
		if err != nil {
			return utils.ErrorsWrap(err, "repository.FindPostByID() error")
		}
		if master.ID != ctxUser.ID {
			return utils.ErrorsNew("Forbbiden to delete the post")
		}
	}

	tx := u.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = u.repository.DeleteTeam(teamID, tx)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeleteTeam() error")
	}
	err = u.repository.DeleteTeamMaster(teamID, tx)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeleteTeamMaster() error")
	}
	err = u.repository.DeleteAllTeamPlayers(teamID, tx)
	if err != nil {
		return utils.ErrorsWrap(err, "repository.DeleteAllTeamPlayers() error")
	}
	return nil
}

// NewUsecase creare new instance of Usecase
func NewUsecase(bu *base.Usecase, db *gorm.DB, r Repository) Usecase {
	return &usecase{*bu, db, r}
}
