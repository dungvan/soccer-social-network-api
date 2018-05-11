package router

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network-api/infrastructure"
	"github.com/dungvan2512/soccer-social-network-api/match"
	"github.com/dungvan2512/soccer-social-network-api/post"
	"github.com/dungvan2512/soccer-social-network-api/shared/base"
	mMiddleware "github.com/dungvan2512/soccer-social-network-api/shared/middleware"
	"github.com/dungvan2512/soccer-social-network-api/team"
	"github.com/dungvan2512/soccer-social-network-api/tournament"
	"github.com/dungvan2512/soccer-social-network-api/user"
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
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {
	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	// Custom middleware(Translation)
	r.Mux.Use(r.TranslationHandler.Middleware.Middleware)
	// Custom middleware(Logger)
	r.Mux.Use(mMiddleware.Logger(r.LoggerHandler))
	// CORS middleware
	r.Mux.Use(mMiddleware.CORS(r.LoggerHandler))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// error handler set.
	eh := base.NewHTTPErrorHandler(r.LoggerHandler.Log)
	r.Mux.NotFound(eh.StatusNotFound)
	r.Mux.MethodNotAllowed(eh.StatusMethodNotAllowed)

	r.Mux.Method(http.MethodGet, "/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// base set.
	bh := base.NewHTTPHandler(r.LoggerHandler.Log)
	// base set.
	br := base.NewRepository(r.LoggerHandler.Log)
	// base set.
	bu := base.NewUsecase(r.LoggerHandler.Log)
	// user set
	uh := user.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)
	// post set
	ph := post.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler, r.S3Handler)
	// team set
	th := team.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)
	// match set
	mh := match.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)
	// match set
	toh := tournament.NewHTTPHandler(bh, bu, br, r.SQLHandler, r.CacheHandler)

	r.Mux.Route("/users", func(cr chi.Router) {
		cr.Post("/register", uh.Register)
		cr.Post("/login", uh.Login)
		cr.With(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB)).
			Get("/{user_name}", uh.Show)
		cr.With(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB)).
			Put("/{id:0*([1-9])([0-9]?)+}", uh.Update)
		cr.With(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB)).
			With(mMiddleware.CheckSuperAdmin(r.LoggerHandler)).
			Get("/", uh.Index)
		cr.With(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB)).
			With(mMiddleware.CheckSuperAdmin(r.LoggerHandler)).
			Delete("/{id:0*([1-9])([0-9]?)+}", uh.Delete)
	})

	r.Mux.Route("/posts", func(cr chi.Router) {
		cr.Use(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB))
		cr.With(mMiddleware.CheckSuperAdmin(r.LoggerHandler)).Get("/", ph.Index)
		cr.Get("/users/{id:0*([1-9])([0-9]?)+}", ph.GetByUserID)
		cr.Post("/", ph.Create)
		cr.Post("/images", ph.UploadImages)
		cr.Route("/{id:0*([1-9])([0-9]?)+}", func(cr chi.Router) {
			cr.Delete("/", ph.Delete)
			cr.Put("/", ph.Update)
			cr.Get("/", ph.Show)
			cr.Post("/star", ph.UpStar)
			cr.Delete("/star", ph.DeleteStar)
			cr.Route("/comment", func(cr chi.Router) {
				cr.Post("/", ph.CommentCreate)
				cr.Route("/{comment_id:0*([1-9])([0-9]?)+}", func(cr chi.Router) {
					cr.Delete("/", ph.CommentDelete)
					cr.Put("/", ph.CommentUpdate)
					cr.Post("/star", ph.CommentUpStar)
					cr.Delete("/star", ph.CommentDeleteStar)
				})
			})
		})
	})

	r.Mux.Route("/teams", func(cr chi.Router) {
		cr.Use(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB))
		cr.With(mMiddleware.CheckSuperAdmin(r.LoggerHandler)).Get("/", th.Index)
		cr.Get("/users/{id:0*([1-9])([0-9]?)+}", th.GetByUser)
		cr.Post("/", th.Create)
		cr.Get("/{id:0*([1-9])([0-9]?)+}", th.Show)
		cr.Delete("/{id:0*([1-9])([0-9]?)+}", th.Delete)
		cr.Put("/{id:0*([1-9])([0-9]?)+}", th.Update)
	})

	r.Mux.Route("/matches", func(cr chi.Router) {
		cr.Use(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB))
		cr.Post("/", mh.Create)
		cr.Get("/{id:0*([1-9])([0-9]?)+}", mh.Show)
	})

	r.Mux.Route("/tournaments", func(cr chi.Router) {
		cr.Use(mMiddleware.JwtAuth(r.LoggerHandler, r.SQLHandler.DB))
		cr.Post("/", toh.Create)
		cr.Get("/{id:0*([1-9])([0-9]?)+}", toh.Show)
	})
}
