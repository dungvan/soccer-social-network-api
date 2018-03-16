package usecase

import "github.com/sirupsen/logrus"

// BaseUsecase struct.
type BaseUsecase struct {
	Logger *logrus.Logger
}

// NewBaseUsecase returns NewBaseUsecase instance.
func NewBaseUsecase(logger *logrus.Logger) *BaseUsecase {
	return &BaseUsecase{Logger: logger}
}
