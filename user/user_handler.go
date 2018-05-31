package user

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/dungvan/soccer-social-network-api/shared/auth"
	"github.com/go-chi/chi"

	"github.com/dungvan/soccer-social-network-api/infrastructure"
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/dungvan/soccer-social-network-api/shared/utils"
	"github.com/sirupsen/logrus"
)

// HTTPHandler interface
type HTTPHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	FriendRequest(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

// handler struct
type handler struct {
	base.HTTPHandler
	usecase Usecase
}

// Register handler
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	i, _ := net.Interfaces()
	fmt.Println(i)
	request := &RegisterReuqest{}
	messages, err := h.ParseJSON(r, request)
	if len(messages) != 0 {
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}
	if err != nil {
		h.Logger.Error(err, "h.ParseJson() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	// call register usecase.
	err = h.usecase.Register(*request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Register() error")
		errString := []string{}
		if strings.Contains(err.Error(), duplidateUniquePreMessage) {
			for _, key := range uniqueKeys {
				if strings.Contains(err.Error(), key) {
					errString = append(errString, fieldName[key]+" alredy exist")
				}
			}
			common := utils.CommonResponse{Message: "Register failed.", Errors: errString}
			h.StatusBadRequest(w, common)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, utils.CommonResponse{Message: "Register successfully.", Errors: nil})
}

// Login handler
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	request := &LoginRequest{}
	messages, err := h.ParseJSON(r, request)
	if len(messages) != 0 {
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}
	if err != nil {
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	token, err := h.usecase.Login(*request)
	if err != nil {
		if err == errUserNameOrPassword {
			h.Logger.WithFields(logrus.Fields{
				"error": err,
			}).Error("usecase.Login() error")
			common := utils.CommonResponse{Message: "Login failed.", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Login() error")
		common := utils.CommonResponse{Message: "Internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}

	h.ResponseJSON(w, token)
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	request := &IndexRequest{}
	h.ParseForm(r, request)
	// validate get data.
	if err := h.Validate(w, request); err != nil {
		return
	}

	response, err := h.usecase.Index(*request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Index() error")
		if response.TypeOfStatusCode == http.StatusBadRequest {
			common := utils.CommonResponse{Message: "Bad request error", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		if response.TypeOfStatusCode == http.StatusNotFound {
			h.StatusNotFoundRequest(w, nil)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

func (h *handler) Show(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")
	response, err := h.usecase.Show(userName)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Show() error")
		if response.TypeOfStatusCode == http.StatusBadRequest {
			common := utils.CommonResponse{Message: "Bad request error response", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		if response.TypeOfStatusCode == http.StatusNotFound {
			h.StatusNotFoundRequest(w, nil)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		common := utils.CommonResponse{Message: "Delete failed", Errors: []string{err.Error()}}
		h.StatusServerError(w, common)
		return
	}

	common := utils.CommonResponse{Message: "Delete success"}
	h.ResponseJSON(w, common)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	request := &UpdateRequest{}
	messages, err := h.ParseJSON(r, request)
	if len(messages) != 0 {
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}
	if err != nil {
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	curUser := auth.GetUserFromContext(r.Context())
	if curUser.ID != request.ID && curUser.Role != "s_admin" {
		common := utils.CommonResponse{Message: "Update failed", Errors: []string{"Forbidden to update this user"}}
		h.StatusBadRequest(w, common)
		return
	}

	user, err := h.usecase.Update(*request)
	if err != nil {
		common := utils.CommonResponse{Message: "Update failed", Errors: []string{err.Error()}}
		h.StatusBadRequest(w, common)
		return
	}

	h.ResponseJSON(w, user)
}

// FriendRequest is user request connection to other
func (h *handler) FriendRequest(w http.ResponseWriter, r *http.Request) {
	request := &FriendRequest{}
	messages, err := h.ParseJSON(r, request)
	if len(messages) != 0 {
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}

	if err != nil {
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	request.UserID = auth.GetUserFromContext(r.Context()).ID

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}
	// call usecase
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) HTTPHandler {
	userRepo := NewRepository(br, s.DB, c.Conn)
	userUsecase := NewUsecase(bu, s.DB, userRepo)
	return &handler{HTTPHandler: *bh, usecase: userUsecase}
}
