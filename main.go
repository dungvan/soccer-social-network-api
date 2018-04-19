package main

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network/infrastructure"
	"github.com/dungvan2512/soccer-social-network/router"
	mMiddleware "github.com/dungvan2512/soccer-social-network/shared/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// start internal server
	go startInt()

	// start external server
	startExt()
}

func startExt() {
	// sql new.
	sqlHandler := infrastructure.NewSQL()
	// s3 new.
	s3Handler := infrastructure.NewS3()
	// cache new.
	cacheHandler := infrastructure.NewCache()
	// logger new.
	loggerHandler := infrastructure.NewLogger()
	// translation new.
	translationHandler := infrastructure.NewTranslation()

	mux := chi.NewRouter()
	r := &router.Router{
		Mux:                mux,
		SQLHandler:         sqlHandler,
		S3Handler:          s3Handler,
		CacheHandler:       cacheHandler,
		LoggerHandler:      loggerHandler,
		TranslationHandler: translationHandler,
	}

	r.InitializeRouter()
	r.SetupHandler()

	// after process
	defer infrastructure.CloseLogger(r.LoggerHandler.Logfile)
	defer infrastructure.CloseRedis(r.CacheHandler.Conn)

	_ = http.ListenAndServe(":8080", mux)
}

func startInt() {
	mux := chi.NewRouter()
	logger := infrastructure.NewLogger()
	mux.Use(mMiddleware.Logger(logger))

	defer infrastructure.CloseLogger(logger.Logfile)

	// profile
	mux.Mount("/debug", middleware.Profiler())
	_ = http.ListenAndServe(":18080", mux)
}
