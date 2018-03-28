package user

import (
	"net/http"

	"github.com/dungvan2512/socker-social-network/shared/base"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// PostRegister handler
func (h *HTTPHandler) PostRegister(w http.ResponseWriter, r *http.Request) {

}
