package base

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Repository struct.
type Repository struct {
	Logger *logrus.Logger
}

// NewRepository returns Repository instance.
func NewRepository(logger *logrus.Logger) *Repository {
	return &Repository{Logger: logger}
}

// GetByte get resp.Body.
func (b *Repository) GetByte(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
