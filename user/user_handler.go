package user

import (
	"net/http"
	"strings"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Register handler
func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	request := &RegisterReuqest{}
	message, err := h.ParseJSON(r, request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Register() error")
		common := utils.CommonResponse{Message: "Parse request error.", Errors: message}
		h.StatusBadRequest(w, common)
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

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	userRepo := NewRepository(br, s.DB, c.Conn)
	userUsecase := NewUsecase(bu, s.DB, userRepo)
	return &HTTPHandler{HTTPHandler: *bh, usecase: userUsecase}
}
