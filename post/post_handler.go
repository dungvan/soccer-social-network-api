package post

import (
	"net/http"
	"strconv"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/auth"
	"github.com/dungvan2512/socker-social-network/shared/base"
	"github.com/dungvan2512/socker-social-network/shared/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserFromContext(r.Context()).ID
	resp, err := h.usecase.Index(userID)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Index() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, resp)
}

// Create a post Handler
func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := CreateRequest{}
	messages, err := h.ParseJSON(r, &request)
	if len(messages) != 0 {
		common := utils.CommonResponse{Message: "validation error.", Errors: messages}
		h.StatusBadRequest(w, common)
		return
	}
	if err != nil {
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	request.User = auth.GetUserFromContext(r.Context())
	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	postID, err := h.usecase.Create(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Create() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, CreateResponse{postID})
}

// UpStar increase the number of "star" about outfit.
//
// "First": to register to OutfitStarCount Outfit ID.
// "Second": to register to PostStarHistory Outfit ID and User ID.
// "Finally": returns latest number of stars about specified outfit to the app.
func (h *HTTPHandler) UpStar(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "post_id"))
	request := StarCountRequest{}
	request.PostID = uint(postID)
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	response, err := h.usecase.CountUpStar(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.CountUpStar() error")
		if response.TypeOfStatusCode == http.StatusBadRequest {
			common := utils.CommonResponse{Message: "Bad request error response", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		if response.TypeOfStatusCode == http.StatusNotFound {
			h.StatusNotFoundRequest(w, nil)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}

	h.ResponseJSON(w, response)
}

// DeleteStar decrease the number of "star" about outfit
//
// "First": rewrites record of table “post_star_history” with key column “id_user_app” and "id_outfit".
// "Second": rewrites record of table “outfit_star_count” with key column "id_outfit".
// "Finally": returns latest number of stars about specified outfit.
func (h *HTTPHandler) DeleteStar(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "post_id"))
	request := StarCountRequest{}
	request.PostID = uint(postID)
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	response, err := h.usecase.CountDownStar(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.CountDownStar() error")
		if response.TypeOfStatusCode == http.StatusBadRequest {
			common := utils.CommonResponse{Message: "Bad request error response", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		if response.TypeOfStatusCode == http.StatusNotFound {
			h.StatusNotFoundRequest(w, nil)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error response", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// UploadImages fot a post
func (h *HTTPHandler) UploadImages(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := UploadImagesRequest{}
	err := h.ParseMultipart(r, request)
	if err != nil {
		common := utils.CommonResponse{Message: "can't receive request parameter.", Errors: nil}
		h.StatusBadRequest(w, common)
		return
	}
	if r.MultipartForm == nil {
		common := utils.CommonResponse{Message: "can't receive request parameter.", Errors: nil}
		h.StatusBadRequest(w, common)
		return
	}
	fileMap := r.MultipartForm.File
	fileHeaders := fileMap["image_files"]
	if fileHeaders != nil {
		request.Images = []Image{}
		for _, fileHeader := range fileHeaders {
			body, err := fileHeader.Open()
			if err != nil {
				if body != nil {
					body.Close()
				}
				h.Logger.Error("can't open file " + fileHeader.Filename)
				continue
			}
			mimeType, err := h.GetFileHeaderContentType(body)
			if err != nil {
				h.Logger.Error("can't detect mime type " + fileHeader.Filename)
				continue
			}
			// roll back cursor to first position in file
			_, err = body.Seek(0, 0)
			if err != nil {
				h.Logger.Error("can't not seek cursor to first position in file.")
			}
			image := Image{Body: body, MimeType: mimeType, Size: fileHeader.Size, Name: fileHeader.Filename}
			request.Images = append(request.Images, image)
		}
	}
	// validate post data.
	if err := h.Validate(w, request); err != nil {
		return
	}
	for index, image := range request.Images {
		if image.Body != nil {
			filename, err := h.GetRandomFileName("", image.Name)
			if err != nil {
				h.Logger.Error("can't get random file name.")
				common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
				h.StatusServerError(w, common)
			}
			request.Images[index].Name = filename
		}
	}
	// Save to S3.
	response, err := h.usecase.UploadImages(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecaseInterface.PostDetectItem() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// UploadVideos fot a post
func (h *HTTPHandler) UploadVideos(w http.ResponseWriter, r *http.Request) {

}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache, s3 *infrastructure.S3) *HTTPHandler {
	// outfit set.
	or := NewRepository(br, s.DB, c.Conn, s3.NewRequest)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
