package post

import (
	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/base"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
