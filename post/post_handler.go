package post

import (
	"net/http"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/auth"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
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
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	// err = h.usecase.PostOutfitDetail(request)
	// if err != nil {
	// 	h.Logger.WithFields(logrus.Fields{
	// 		"error": err,
	// 	}).Error("usecaseInterface.PostOutfitDetail() error")
	// 	common := CommonResponse{Message: "internal server error.", Errors: nil}
	// 	h.StatusServerError(w, common)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	return
}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
