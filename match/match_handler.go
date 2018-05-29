package match

import (
	"net/http"
	"strconv"

	"github.com/dungvan/soccer-social-network-api/infrastructure"
	"github.com/dungvan/soccer-social-network-api/shared/auth"
	"github.com/dungvan/soccer-social-network-api/shared/base"
	"github.com/dungvan/soccer-social-network-api/shared/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Create handler
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	// maping request
	request := CreateRequest{}
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	messages, err := h.ParseJSON(r, &request)
	if len(messages) != 0 {
		h.Logger.Error(err, "h.ParseJSON() error")
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}
	if err != nil {
		h.Logger.Error(err, "h.ParseJSON() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	matchID, err := h.usecase.Create(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Create() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, CreateResponse{matchID})
}

// Show handler
func (h *HTTPHandler) Show(w http.ResponseWriter, r *http.Request) {
	matchID, err := strconv.Atoi(chi.URLParam(r, "id"))
	response, err := h.usecase.Show(uint(matchID))
	if response.TypeOfStatusCode == http.StatusNotFound {
		h.StatusNotFoundRequest(w, nil)
		return
	}
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Show() error")
		common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}

	h.ResponseJSON(w, response)
}

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	request := &IndexRequest{}
	h.ParseForm(r, request)
	// validate get data.
	if err := h.Validate(w, request); err != nil {
		return
	}
	response, err := h.usecase.Index(request.Page)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Index() error")
		common := utils.CommonResponse{Message: "Internal server error", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// GetByUserName handler
func (h *HTTPHandler) GetByUserName(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "user_name")

	resp, err := h.usecase.GetByUserName(userName)
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

// GetByMaster handler
func (h *HTTPHandler) GetByMaster(w http.ResponseWriter, r *http.Request) {
	id := auth.GetUserFromContext(r.Context()).ID
	resp, err := h.usecase.GetByMasterUserID(id)
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

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
