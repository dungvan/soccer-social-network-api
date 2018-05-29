package utils

import (
	"io"
	"io/ioutil"
	"os"

	detection "github.com/dungvan/soccer-social-network-api/shared/file-detection"
)

// ExistsFile check exists file.
func ExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// CopyFile copy file.
func CopyFile(srcPath, dstPath string) (err error) {
	src, err := os.Open(srcPath)
	if err != nil {
		return ErrorsWrap(err, "can't open file")
	}
	defer func() {
		if srcerr := src.Close(); err == nil && srcerr != nil {
			err = ErrorsWrap(srcerr, "can't close file")
		}
	}()
	dst, err := os.Create(dstPath)
	if err != nil {
		return ErrorsWrap(err, "can't create file")
	}
	defer func() {
		if dsterr := dst.Close(); err == nil && dsterr != nil {
			err = ErrorsWrap(dsterr, "can't close file")
		}
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		return ErrorsWrap(err, "can't copy file")
	}
	return err
}

// ReadByteFile get file bytes data.
func ReadByteFile(filepath string) (b []byte, err error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, ErrorsWrap(err, "can't open file")
	}
	defer func() {
		if fcerr := f.Close(); err == nil && fcerr != nil {
			err = ErrorsWrap(fcerr, "can't close file")
		}
	}()

	fb, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, ErrorsWrap(err, "can't read file")
	}
	return fb, err
}

// DetectFileContentType return the content type of file by 128 first bytes
func DetectFileContentType(data []byte) string {
	return detection.DetectContentType(data)
}
