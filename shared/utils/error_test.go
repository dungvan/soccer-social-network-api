package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsNew(t *testing.T) {
	err := ErrorsNew("test error")
	assert.Error(t, err)
	assert.Regexp(t, "test error", err)
}
func TestErrorsWrap(t *testing.T) {
	err := ErrorsNew("test error1")
	err = ErrorsWrap(err, "test error2")
	assert.Error(t, err)
	assert.Regexp(t, "test error1", err)
	assert.Regexp(t, "test error2", err)
}

func TestErrorsWrapf(t *testing.T) {
	err := ErrorsNew("test error1")
	err = ErrorsWrapf(err, "test error%v", "2")
	assert.Error(t, err)
	assert.Regexp(t, "test error1", err)
	assert.Regexp(t, "test error2", err)
}
