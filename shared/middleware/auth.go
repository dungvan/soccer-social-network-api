package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dungvan2512/soccer-social-network/model"
	"github.com/dungvan2512/soccer-social-network/shared/auth"
	"github.com/jinzhu/gorm"

	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/shared/utils"
	"github.com/sirupsen/logrus"
)

// AuthFailedResponse struct.
type AuthFailedResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// DefaultUnauthorizedResponse function return AuthFailedResponse
func DefaultUnauthorizedResponse() AuthFailedResponse {
	return AuthFailedResponse{Message: "Unauthorized", Errors: []string{}}
}

// JwtAuth check jwt header parse and validation.
func JwtAuth(logger *infrastructure.Logger, db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtHeader := infrastructure.GetConfigString("jwt.header")
			tokenString := r.Header.Get(jwtHeader)
			// Check token string with format
			checkFormat := strings.HasPrefix(tokenString, "Bearer ")
			if checkFormat != true {
				utils.ResponseJSON(w, http.StatusUnauthorized, DefaultUnauthorizedResponse())
				return
			}
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			user, err := auth.ParseToken(tokenString)
			if err != nil {
				logger.Log.WithFields(logrus.Fields{"error": err}).Info("parse token fail")
				utils.ResponseJSON(w, http.StatusUnauthorized, DefaultUnauthorizedResponse())
				return
			}
			userModel := model.User{}
			err = db.Where(`id=?`, user.ID).Where("email = ?", user.Email).First(&userModel).Error
			if err != nil {
				logger.Log.WithFields(logrus.Fields{"error": err}).Info("user dose'nt exist")
				utils.ResponseJSON(w, http.StatusUnauthorized, DefaultUnauthorizedResponse())
				return
			}
			ctx := context.WithValue(r.Context(), auth.ContextKeyAuth, userModel)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// CheckSuperAdmin middleware
func CheckSuperAdmin(logger *infrastructure.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role := auth.GetUserFromContext(r.Context()).Role
			if role != "s_admin" {
				logger.Log.WithFields(logrus.Fields{"error": errors.New("role of user is " + role)}).Info("user are not super admin")
				utils.ResponseJSON(w, http.StatusUnauthorized, DefaultUnauthorizedResponse())
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
