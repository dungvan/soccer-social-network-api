package base

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name string `form:"name" json:"name" validate:"required"`
	Pass string `form:"pass" json:"pass" validate:"required"`
}
type ErrorUser struct {
	Name int `form:"name" json:"name" validate:"required"`
	Pass int `form:"pass" json:"pass" validate:"required"`
}
type Image struct {
	FileType string `form:"file_type" validate:"omitempty,eq=image/bmp|eq=image/dib|eq=image/jpeg|eq=image/jp2|eq=image/png|eq=image/webp|eq=image/x-portable-anymap|eq=image/x-portable-bitmap|eq=image/x-portable-graymap|eq=image/x-portable-pixmap|eq=image/x-cmu-raster|eq=image/tiff|eq=image/gif"`
}

var fileContent = []string{"image/bmp", "image/dib", "image/jpeg", "image/jp2", "image/png", "image/webp", "image/x-portable-anymap", "image/x-portable-bitmap", "image/x-portable-graymap", "image/x-portable-pixmap", "image/x-cmu-raster", "image/tiff", "image/gif"}

func TestHandlerNewHTTPHandler(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)
	assert.NotEmpty(t, bh)
}

func TestHandlerParse(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	// request param
	values := url.Values{
		"name": {"circle"},
		"pass": {"pass"},
	}
	body := bytes.NewBufferString(values.Encode())
	// request
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	user := User{}
	err := bh.Parse(req, &user)
	assert.Equal(t, "circle", user.Name)
	assert.Equal(t, "pass", user.Pass)
	assert.NoError(t, err)

	// r is nil.
	err = bh.Parse(nil, &user)
	assert.Error(t, err)
	// data is nil.
	err = bh.Parse(req, nil)
	assert.Error(t, err)
	// error struct.
	errorUser := ErrorUser{}
	err = bh.Parse(req, &errorUser)
	assert.Error(t, err)
}

func TestHandlerParseForm(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)
	// request
	req, _ := http.NewRequest("GET", "/?name=circle&pass=pass", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	user := User{}
	bh.ParseForm(req, &user)
	assert.Equal(t, "circle", user.Name)

}

func TestHandlerMultipartParse(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	w.WriteField("name", "circle")
	w.WriteField("pass", "pass")
	w.Close()
	req, _ := http.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", w.FormDataContentType())

	user := User{}
	err := bh.ParseMultipart(req, &user)
	assert.Equal(t, "circle", user.Name)
	assert.Equal(t, "pass", user.Pass)
	assert.NoError(t, err)

	// req is nil.
	err = bh.ParseMultipart(nil, &user)
	assert.Error(t, err)
	// user is nil.
	err = bh.ParseMultipart(req, nil)
	assert.Error(t, err)
}

func TestHandlerSaveToFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		bh.SaveToFile(req, "text_file", "dst_sample.txt")

		buf, _ := ioutil.ReadFile("dst_sample.txt")

		assert.Equal(t, "sample text", string(buf))

		os.Remove("dst_sample.txt")
	})
	t.Run("failed by nothing formfile", func(t *testing.T) {
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		req, _ := http.NewRequest("POST", "/", nil)

		err := bh.SaveToFile(req, "", "dst_sample.txt")

		assert.Error(t, err)
	})
	t.Run("failed by nothing tofile", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		err := bh.SaveToFile(req, "text_file", "")

		assert.Error(t, err)
	})
	t.Run("failed by nil", func(t *testing.T) {
		sampleFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/data/sample.txt"

		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		body := new(bytes.Buffer)
		w := multipart.NewWriter(body)
		utils.MultipartFileWriter(w, "text_file", sampleFile)
		w.Close()
		req, _ := http.NewRequest("POST", "/", body)
		req.Header.Set("Content-Type", w.FormDataContentType())

		err := bh.SaveToFile(nil, "text_file", "")

		assert.Error(t, err)
	})
}

func TestHandlerGetRandomTempFileName(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)
	name, err := bh.GetRandomTempFileName("prefix-", "sample.txt")
	assert.Regexp(t, "prefix-", name)
	assert.NoError(t, err)
}

func TestHandlerGetRandomFileName(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)
	fileName, err := bh.GetRandomFileName("file-", "example.jpg")
	assert.Regexp(t, "file-", fileName)
	assert.NoError(t, err)
}

func TestHandlerResponseJson(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	// server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.ResponseJSON(w, user)
		return
	})
	// httpTest newServer
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// client
	res, err := http.Get(ts.URL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	resUser := User{}
	json.Unmarshal(buf, &resUser)
	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.Equal(t, user.Name, resUser.Name)
	assert.Equal(t, user.Pass, resUser.Pass)
}

func TestHandlerStatusRedirect(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.StatusRedirect(w, "/redirect")
		return
	})

	// router
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	mux.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusFound, rec.Code)
	path := rec.Header().Get("Location")
	assert.Equal(t, "/redirect", path)
}

func TestHandlerStatusBadRequest(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bh.StatusBadRequest(w, user)
		return
	})

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	handler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	resUser := User{}
	json.Unmarshal(rec.Body.Bytes(), &resUser)
	assert.NotEmpty(t, user)
	assert.Equal(t, user.Name, resUser.Name)
	assert.Equal(t, user.Pass, resUser.Pass)
}

func TestHandlerStatusNotFoundRequest(t *testing.T) {
	user := User{
		"circle",
		"password",
	}

	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	t.Run("response json", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusNotFoundRequest(w, user)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		resUser := User{}
		json.Unmarshal(rec.Body.Bytes(), &resUser)
		assert.Equal(t, user.Name, resUser.Name)
		assert.Equal(t, user.Pass, resUser.Pass)
	})
	t.Run("no response json.", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusNotFoundRequest(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		assert.Nil(t, rec.Body.Bytes())
	})
}

func TestHandlerStatusServerError(t *testing.T) {
	t.Run("status server error", func(t *testing.T) {
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusServerError(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("status server error with body", func(t *testing.T) {
		user := User{
			"circle",
			"password",
		}

		bh := NewHTTPHandler(infrastructure.NewLogger().Log)

		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusServerError(w, user)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		resUser := User{}
		json.Unmarshal(rec.Body.Bytes(), &resUser)
		assert.Equal(t, user.Name, resUser.Name)
		assert.Equal(t, user.Pass, resUser.Pass)
	})
}

func TestHandlerGetTranslaterFunc(t *testing.T) {
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	translation := infrastructure.NewTranslation()

	// server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translaterFunc := bh.GetTranslaterFunc(r)
		assert.NotEmpty(t, translaterFunc)
		return
	})
	// router
	mux := chi.NewRouter()
	mux.Use(translation.Middleware.Middleware)
	mux.HandleFunc("/", handler)

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	mux.ServeHTTP(rec, req)
}

func TestHandlerGetFileHeaderContentType(t *testing.T) {
	t.Run("get file header content type success", func(t *testing.T) {
		imageFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/image/W_ULD_171201_01.jpg"

		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		f, _ := os.Open(imageFile)
		defer f.Close()
		fileContentType, _ := bh.GetFileHeaderContentType(f)
		assert.Equal(t, "image/jpeg", fileContentType)
	})
	t.Run("arguments is nil", func(t *testing.T) {
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		f, err := bh.GetFileHeaderContentType(nil)
		assert.Error(t, err)
		assert.Empty(t, f)
	})
}
func TestHandlerValidatorFunc(t *testing.T) {
	t.Run("validate success", func(t *testing.T) {
		user := User{
			Name: "username",
			Pass: "pass",
		}
		// base handler
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		// response write
		w := httptest.NewRecorder()

		err := bh.Validate(w, user)
		assert.NoError(t, err)
	})
	t.Run("validate faild", func(t *testing.T) {
		user := User{
			Pass: "password",
		}
		// base handler
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		// response writer
		w := httptest.NewRecorder()

		err := bh.Validate(w, user)
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("validate faild with file text/css", func(t *testing.T) {
		image := Image{
			FileType: "text/css",
		}
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		w := httptest.NewRecorder()

		err := bh.Validate(w, image)
		assert.Error(t, err)
		assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
	})
	for _, v := range fileContent {
		t.Run("validate success with image valid "+v, func(t *testing.T) {
			image := Image{
				FileType: v,
			}
			bh := NewHTTPHandler(infrastructure.NewLogger().Log)
			w := httptest.NewRecorder()
			err := bh.Validate(w, image)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, w.Code)
		})
	}
}

func TestHandlerStatusUnsupportedMediaTypeRequest(t *testing.T) {
	image := Image{
		FileType: "text/css",
	}
	bh := NewHTTPHandler(infrastructure.NewLogger().Log)

	t.Run("response json", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusUnsupportedMediaTypeRequest(w, image)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)

		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		resImage := Image{}
		json.Unmarshal(rec.Body.Bytes(), &resImage)
		assert.Equal(t, image.FileType, resImage.FileType)
	})
	t.Run("no response json.", func(t *testing.T) {
		// handler
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bh.StatusUnsupportedMediaTypeRequest(w, nil)
			return
		})

		// response writer
		rec := httptest.NewRecorder()
		// new request
		req, _ := http.NewRequest("GET", "/", nil)
		// request
		handler(rec, req)
		assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
		assert.Nil(t, rec.Body.Bytes())
	})
}
func TestParseJSONSuccess(t *testing.T) {
	testCases := []struct {
		description string
		dataTest    []byte
		structTest  interface{}
	}{
		{
			description: "simple struct!",
			dataTest: []byte(`{
					"caption": "caption",
					"source_image_file_name": "2",
					"place_id": "place id"
				}`),
			structTest: struct {
				Caption             string `json:"caption"`
				SourceImageFileName string `json:"source_image_file_name"`
				PlaceID             string `json:"place_id"`
			}{},
		},
		{
			description: "struct has child struct!",
			dataTest: []byte(`{
					"caption": "caption",
					"place_id": "place id",
					"item":{
						"im_name": "d",
						"product_name": "Tシャツ"
						}
				}`),
			structTest: struct {
				Caption string `json:"caption"`
				PlaceID string `json:"place_id"`
				Item    struct {
					ImageName   string `json:"im_name"`
					ProductName string `json:"product_name"`
				} `json:"item"`
			}{},
		},
		{
			description: "struct has array!",
			dataTest: []byte(`{
					"caption": "caption",
					"place_id": "place id",
					"item":[
						"im_name",
						"product_name"
					]
				}`),
			structTest: struct {
				Caption string   `json:"caption"`
				PlaceID string   `json:"place_id"`
				Item    []string `json:"item"`
			}{},
		},
		{
			description: "struct child has array!",
			dataTest: []byte(`{
					"caption": "caption",
					"place_id": "place id",
					"item":{
						"im_name":"image name",
						"img_box":[1,2,3,4]
					}
				}`),
			structTest: struct {
				Caption string `json:"caption"`
				PlaceID string `json:"place_id"`
				Item    struct {
					ImageName string `json:"im_name"`
					ImageBox  []uint `json:"img_box"`
				} `json:"item"`
			}{},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			bh := NewHTTPHandler(infrastructure.NewLogger().Log)
			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(testCase.dataTest))
			req.Header.Set("Content-Type", "application/json")
			messages, err := bh.ParseJSON(req, &testCase.structTest)
			assert.NoError(t, err)
			assert.Nil(t, messages)
		})
	}
}
func TestParseJSONFail(t *testing.T) {

	type StructTest struct {
		Caption             string `json:"caption"`
		SourceImageFileName string `json:"source_image_file_name"`
		PlaceID             string `json:"place_id"`
	}
	testCases := []struct {
		description string
		dataTest    []byte
		errExpect   error
		messExpect  []string
	}{
		{
			description: "fail only one type of simple struct",
			dataTest: []byte(`{
					"caption": 1,
					"source_image_file_name": "2",
					"place_id": "place id"
				}`),
			errExpect:  errors.New("can't decoder.Decode error"),
			messExpect: []string{"caption is error."},
		},
		{
			description: "fail more type of simple struct",
			dataTest: []byte(`{
					"caption": 1,
					"source_image_file_name": "2",
					"place_id": 1
				}`),
			errExpect:  errors.New("can't decoder.Decode error"),
			messExpect: []string{"caption is error.", "place_id is error."},
		},
		{
			description: "fail format json",
			dataTest: []byte(`{
						"caption": {1},
						"source_image_file_name": "2",
						"place_id": 1
					}`),
			errExpect:  errors.New("JSON format is error"),
			messExpect: []string{"JSON input invalid."},
		},
		{
			description: "fail read from request",
			dataTest:    nil,
			errExpect:   errors.New("JSON format is error"),
			messExpect:  []string{"JSON input invalid."},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			bh := NewHTTPHandler(infrastructure.NewLogger().Log)
			req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(testCase.dataTest))
			req.Header.Set("Content-Type", "application/json")
			messages, err := bh.ParseJSON(req, &StructTest{})
			assert.Contains(t, err.Error(), testCase.errExpect.Error())
			assert.Equal(t, testCase.messExpect, messages)
		})
	}
}

func TestIsJSONFunc(t *testing.T) {
	testCase := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "Test isJSON return true when string json valid",
			input: `{
				"outfit": []
			  }`,
			expected: true,
		},
		{
			name:     "Test isJON return false when string input is not json",
			input:    `is not json`,
			expected: false,
		},
	}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			actual := isJSON(v.input)
			assert.Equal(t, v.expected, actual)
		})
	}
}

func TestHTTPHandlerDetectImageDimention(t *testing.T) {
	t.Run("get image file dimention success", func(t *testing.T) {
		imageFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/image/W_ULD_171201_01.jpg"
		f, err := os.Open(imageFile)
		if err != nil {
			t.Errorf("could not open image file")
		}
		defer f.Close()
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		width, height, err := bh.DetectImageDimention(f)
		assert.Nil(t, err)
		assert.True(t, width > 0)
		assert.True(t, height > 0)
	})
	t.Run("arguments is not support type", func(t *testing.T) {
		imageFile := os.Getenv("FR_CIRCLE_API_DIR") + "/test/txt/FILE_NOT_IMAGE.txt"
		f, err := os.Open(imageFile)
		if err != nil {
			t.Errorf("could not open image file")
		}
		defer f.Close()
		bh := NewHTTPHandler(infrastructure.NewLogger().Log)
		width, height, err := bh.DetectImageDimention(f)
		assert.Error(t, err)
		assert.Equal(t, 0, width)
		assert.Equal(t, 0, height)
	})
}
