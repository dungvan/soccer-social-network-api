package router

import (
	"net/http"

	"github.com/dungvan2512/socker-social-network/infrastructure"
	module "github.com/dungvan2512/socker-social-network/sample-module"
	"github.com/dungvan2512/socker-social-network/shared/base"
	mMiddleware "github.com/dungvan2512/socker-social-network/shared/middleware"
	"github.com/fastretailing1/circle-api/shared/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux                *chi.Mux
	SQLHandler         *infrastructure.SQL
	S3Handler          *infrastructure.S3
	CacheHandler       *infrastructure.Cache
	LoggerHandler      *infrastructure.Logger
	TranslationHandler *infrastructure.Translation
	SearchAPIHandler   infrastructure.SearchAPI
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {

	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := handler.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// base set.
	bh := base.NewHTTPHandler(r.LoggerHandler.Log)
	// base set.
	br := base.NewRepository(r.LoggerHandler.Log)
	// base set.
	bu := base.NewUsecase(r.LoggerHandler.Log)
	// outfit set.
	mh := module.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler, r.S3Handler)
	r.Mux.Route("/sample", func(cr chi.Router) {
		cr.Get("/", mh.SampleHandler)
	})
}
