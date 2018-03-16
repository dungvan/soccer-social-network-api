package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dungvan2512/socker-social-network/shared/auth"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/dungvan2512/socker-social-network/shared/utils"
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
func JwtAuth(logger *infrastructure.Logger) func(http.Handler) http.Handler {
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
			ctx := context.WithValue(r.Context(), auth.ContextKeyAuth, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
