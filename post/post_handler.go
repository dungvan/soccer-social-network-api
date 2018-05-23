package post

import (
	"net/http"
	"strconv"

	"github.com/dungvan2512/soccer-social-network-api/infrastructure"
	"github.com/dungvan2512/soccer-social-network-api/shared/auth"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	"github.com/dungvan2512/soccer-social-network-api/shared/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	base.HTTPHandler
	usecase Usecase
}

//===========================================
//====================POST===================
//===========================================

// Index handler
func (h *HTTPHandler) Index(w http.ResponseWriter, r *http.Request) {
	request := &IndexRequest{}
	userID := auth.GetUserFromContext(r.Context()).ID
	h.ParseForm(r, request)
	// validate get data.
	if err := h.Validate(w, request); err != nil {
		return
	}

	response, err := h.usecase.Index(userID, request.Page)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Index() error")
		if response.TypeOfStatusCode == http.StatusBadRequest {
			common := utils.CommonResponse{Message: "Bad request error", Errors: []string{err.Error()}}
			h.StatusBadRequest(w, common)
			return
		}
		if response.TypeOfStatusCode == http.StatusNotFound {
			h.StatusNotFoundRequest(w, nil)
			return
		}
		common := utils.CommonResponse{Message: "Internal server error", Errors: []string{}}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, response)
}

// GetByUserID handler
func (h *HTTPHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	request := &IndexRequest{}
	h.ParseForm(r, request)
	// validate get data.
	if err := h.Validate(w, request); err != nil {
		return
	}
	userIDCreate, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userIDCall := auth.GetUserFromContext(r.Context()).ID
	resp, err := h.usecase.GetByUserID(uint(userIDCreate), userIDCall, request.Page)
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
	request.UserID = auth.GetUserFromContext(r.Context()).ID
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

// Show a post Handler
func (h *HTTPHandler) Show(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "id"))
	response, err := h.usecase.Show(uint(postID))
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

// UpStar increase the number of "star" about post.
//
// "First": to register to PostStarCount Post ID.
// "Second": to register to PostStarHistory Post ID and User ID.
// "Finally": returns latest number of stars about specified post to the app.
func (h *HTTPHandler) UpStar(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := StarCountRequest{}
	request.ID = uint(postID)
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

// DeleteStar decrease the number of "star" about post
//
// "First": rewrites record of table “post_star_history” with key column “id_user_app” and "id_post".
// "Second": rewrites record of table “post_star_count” with key column "id_post".
// "Finally": returns latest number of stars about specified post.
func (h *HTTPHandler) DeleteStar(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "id"))
	request := StarCountRequest{}
	request.ID = uint(postID)
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

// Update handler
func (h *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	request := &UpdateRequest{}
	request.ID = uint(postID)
	messages, err := h.ParseJSON(r, request)
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

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}
	curUser := auth.GetUserFromContext(r.Context())
	post, err := h.usecase.Update(*request, curUser)

	if err != nil {
		common := utils.CommonResponse{Message: "Update failed", Errors: []string{err.Error()}}
		h.StatusBadRequest(w, common)
		return
	}

	h.ResponseJSON(w, post)
}

// Delete handler
func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	curUser := auth.GetUserFromContext(r.Context())
	if err := h.usecase.Delete(uint(id), curUser); err != nil {
		common := utils.CommonResponse{Message: "Delete failed", Errors: []string{err.Error()}}
		h.StatusServerError(w, common)
		return
	}

	common := utils.CommonResponse{Message: "Delete success"}
	h.ResponseJSON(w, common)
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
	fileHeaders := fileMap["imageFiles"]

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

	// validate post data.
	if err := h.Validate(w, request); err != nil {
		return
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
	h.ResponseJSON(w, response.ImageNames)
}

// UploadVideos fot a post
func (h *HTTPHandler) UploadVideos(w http.ResponseWriter, r *http.Request) {

}

//===========================================
//==================COMMENT==================
//===========================================

// CommentCreate a comment Handler
func (h *HTTPHandler) CommentCreate(w http.ResponseWriter, r *http.Request) {
	// mapping comment to struct.
	request := CreateCommentRequest{}
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

	postID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	userID := auth.GetUserFromContext(r.Context()).ID
	request.PostID = uint(postID)
	request.UserID = userID
	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}

	resp, err := h.usecase.CommentCreate(request)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("usecase.Create() error")
		common := utils.CommonResponse{Message: "internal server error.", Errors: nil}
		h.StatusServerError(w, common)
		return
	}
	h.ResponseJSON(w, resp)
}

// CommentUpStar increase the number of "star" about post.
//
// "First": to register to PostStarCount Post ID.
// "Second": to register to CommentStarHistory Post ID and User ID.
// "Finally": returns latest number of stars about specified post to the app.
func (h *HTTPHandler) CommentUpStar(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "comment_id"))
	request := StarCountRequest{}
	request.ID = uint(commentID)
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	response, err := h.usecase.CommentCountUpStar(request)
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

// CommentDeleteStar decrease the number of "star" about post
//
// "First": rewrites record of table “comment_star_history” with key column “id_user_app” and "id_post".
// "Second": rewrites record of table “post_star_count” with key column "id_post".
// "Finally": returns latest number of stars about specified post.
func (h *HTTPHandler) CommentDeleteStar(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "comment_id"))
	request := StarCountRequest{}
	request.ID = uint(commentID)
	request.UserID = auth.GetUserFromContext(r.Context()).ID
	response, err := h.usecase.CommentCountDownStar(request)
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

// CommentUpdate handler
func (h *HTTPHandler) CommentUpdate(w http.ResponseWriter, r *http.Request) {
	commentID, _ := strconv.Atoi(chi.URLParam(r, "comment_id"))
	request := &UpdateCommentRequest{}
	request.ID = uint(commentID)
	messages, err := h.ParseJSON(r, request)
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

	// validate get data.
	if err = h.Validate(w, request); err != nil {
		return
	}
	curUser := auth.GetUserFromContext(r.Context())
	err = h.usecase.CommentUpdate(*request, curUser)

	if err != nil {
		common := utils.CommonResponse{Message: "Update failed", Errors: []string{err.Error()}}
		h.StatusBadRequest(w, common)
		return
	}

	h.ResponseJSON(w, utils.CommonResponse{Message: "Update success"})
}

// CommentDelete handler
func (h *HTTPHandler) CommentDelete(w http.ResponseWriter, r *http.Request) {
	commentID, _ := strconv.Atoi(chi.URLParam(r, "comment_id"))
	curUser := auth.GetUserFromContext(r.Context())
	if err := h.usecase.CommentDelete(uint(commentID), curUser); err != nil {
		common := utils.CommonResponse{Message: "Delete failed", Errors: []string{err.Error()}}
		h.StatusServerError(w, common)
		return
	}
	common := utils.CommonResponse{Message: "Delete success"}
	h.ResponseJSON(w, common)
}

// NewHTTPHandler return new HTTPHandler instance.
func NewHTTPHandler(bh *base.HTTPHandler, bu *base.Usecase, br *base.Repository, s *infrastructure.SQL, c *infrastructure.Cache, s3 *infrastructure.S3) *HTTPHandler {
	// post set.
	or := NewRepository(br, s.DB, c.Conn, s3.NewRequest)
	ou := NewUsecase(bu, s.DB, or)
	return &HTTPHandler{*bh, ou}
}
