package post

import (
	"net/http"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/auth"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	user := auth.GetUserFromContext(r.Context())
	resp, err := h.usecase.Index(user.ID)
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

// Create a post Handler
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := CreateRequest{}
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
	request.User = auth.GetUserFromContext(r.Context())
	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	postID, err := h.usecase.Create(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Create() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, CreateResponse{postID})
}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
