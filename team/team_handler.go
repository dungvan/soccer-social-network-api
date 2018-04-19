package team

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/shared/base"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
}

// Create handler
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
}

// Update handler
func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
}

// Show handler
func (h *HTTPHandler) Show(w http.ResponseWriter, r *http.Request) {
	h.ResponseJSON(w, struct {
		Message string `json:"message"`
	}{"this is sample response"})
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
