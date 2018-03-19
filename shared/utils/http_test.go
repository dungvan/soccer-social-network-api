package utils

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultipartFileWriterOK(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	imageFilePath := dir + "/test/image/W_ULD_171201_01.jpg"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := MultipartFileWriter(writer, "image_file", imageFilePath)
	assert.NoError(t, err)

}

func TestMultipartFileWriterNGFileOpen(t *testing.T) {
	dir := os.Getenv("FR_CIRCLE_API_DIR")
	imageFilePath := dir + "/notfoundfile.txt"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	err := MultipartFileWriter(writer, "image_file", imageFilePath)
	assert.Error(t, err)
}

func TestResponseJSONHasData(t *testing.T) {
	type dataTest struct {
		Message string
		Number  int
	}
	w := httptest.NewRecorder()
	data := &dataTest{
		Message: "test",
		Number:  10,
	}
	ResponseJSON(w, http.StatusOK, data)
	assert.Equal(t, http.StatusOK, w.Code)
	var output dataTest
	json.Unmarshal(w.Body.Bytes(), &output)
	assert.Equal(t, data.Message, output.Message)
	assert.Equal(t, data.Number, output.Number)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
}

func TestResponseJSONwithEmptyData(t *testing.T) {
	type dataTest struct {
		Message string
		Number  int
	}
	w := httptest.NewRecorder()
	ResponseJSON(w, http.StatusOK, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	var output dataTest
	json.Unmarshal(w.Body.Bytes(), &output)
	assert.Empty(t, output.Message)
	assert.Empty(t, output.Number)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
}
