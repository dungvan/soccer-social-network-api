package middleware

import (
	"net/http"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	"github.com/sirupsen/logrus"
)

// Logger is log middleware.
func Logger(logger *infrastructure.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger.Log.Info("----------------------------------------------")
			// output headers.
			for i, v := range r.Header {
				logger.Log.WithFields(logrus.Fields{
					"header": i,
					"value":  v,
				}).Info("")
			}
			logger.Log.Info("----------------------------------------------")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Header is log middleware.
func Header(logger *infrastructure.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
