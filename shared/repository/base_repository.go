package repository

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// BaseRepository struct.
type BaseRepository struct {
	Logger *logrus.Logger
}

// NewBaseRepository returns NewBaseRepository instance.
func NewBaseRepository(logger *logrus.Logger) *BaseRepository {
	return &BaseRepository{Logger: logger}
}

// GetByte get resp.Body.
func (b *BaseRepository) GetByte(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}
