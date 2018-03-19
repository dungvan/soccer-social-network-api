package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

// MultipartFileWriter write file and file headers.
func MultipartFileWriter(writer *multipart.Writer, paramName, path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return ErrorsWrapf(err, "can't open error (%v)", path)
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return ErrorsWrapf(err, "can't read file (%v)", path)
	}
	fi, err := file.Stat()
	if err != nil {
		return ErrorsWrapf(err, "can't stat file (%v)", path)
	}
	defer func() {
		if ferr := file.Close(); err == nil && ferr != nil {
			err = ErrorsWrap(err, "can't close file")
		}
	}()

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			quoteEscaper.Replace(paramName), quoteEscaper.Replace(fi.Name())))
	h.Set("Content-Type", DetectFileContentType(fileContents))
	part, err := writer.CreatePart(h)
	if err != nil {
		return ErrorsWrapf(err, "can't create part (%v)", path)
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return ErrorsWrapf(err, "can't write part (%v)", path)
	}
	return err
}

// ResponseJSON function response as json with ResponseWriter
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}
