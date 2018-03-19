package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// ErrorsNew add caller/line and set errors.New.
func ErrorsNew(s string) error {
	path := strings.Trim(getCallerString(), os.Getenv("GOPATH"))
	return errors.New(path + s)
}

// ErrorsWrap add caller/line and set errors.Wrap.
func ErrorsWrap(err error, s string) error {
	path := strings.Trim(getCallerString(), os.Getenv("GOPATH"))
	return errors.Wrap(err, path+s)
}

// ErrorsWrapf add caller/line and set errors.Wrapf.
func ErrorsWrapf(err error, s string, args interface{}) error {
	path := strings.Trim(getCallerString(), os.Getenv("GOPATH"))
	return errors.Wrapf(err, path+s, args)
}

// GetCallerString get caller string.
func getCallerString() string {
	_, f, l, _ := runtime.Caller(2)
	return fmt.Sprintf("%v:%v:", f, l)
}
