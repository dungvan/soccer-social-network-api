package team

import (
	"net/http"
	"strconv"

	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/shared/auth"
	"github.com/dungvan2512/soccer-social-network/shared/base"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserFromContext(r.Context()).ID
	resp, err := h.usecase.Index(userID)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Index() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, resp)
}

// Create handler
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	// maping request
	request := CreateRequest{}
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	messages, err := h.ParseJSON(r, &request)
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

	teamID, err := h.usecase.Create(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Create() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, CreateResponse{teamID})
}

// Update handler
func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
}

// Show handler
func (h *HTTPHandler) Show(w http.ResponseWriter, r *http.Request) {
	teamID, err := strconv.Atoi(chi.URLParam(r, "team_id"))
	response, err := h.usecase.Show(uint(teamID))
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.CountUpStar() error")
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

// Delete handler
func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
