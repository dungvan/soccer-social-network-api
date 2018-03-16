package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPErrorHandler base handler struct.
type HTTPErrorHandler struct {
	BaseHTTPHandler
	logger *logrus.Logger
}

// StatusNotFound responses status code 404.
func (h *HTTPErrorHandler) StatusNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// status code 404
	w.WriteHeader(http.StatusNotFound)
}

// StatusMethodNotAllowed responses status code 405.
func (h *HTTPErrorHandler) StatusMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// status code 405
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// NewHTTPErrorHandler responses new HTTPArticleHandler instance.
func NewHTTPErrorHandler(logger *logrus.Logger) *HTTPErrorHandler {
	return &HTTPErrorHandler{logger: logger}
}
