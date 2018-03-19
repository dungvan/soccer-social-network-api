package base

import "github.com/sirupsen/logrus"

// Usecase struct.
type Usecase struct {
	Logger *logrus.Logger
}

// NewUsecase returns NewBaseUsecase instance.
func NewUsecase(logger *logrus.Logger) *Usecase {
	return &Usecase{Logger: logger}
}
