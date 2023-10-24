package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jimbo8702/lets_get_one/db"
	"github.com/Jimbo8702/lets_get_one/render"
	"github.com/Jimbo8702/lets_get_one/validator"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sagikazarmark/slog-shim"
)

type API struct {
	store 		*db.Store
	log 		*slog.Logger
	validator 	validator.Validator
	session 	*scs.SessionManager
	router 		*chi.Mux
	engine 		render.Renderer
}

func New(db *db.Store, ls *slog.Logger, vd validator.Validator, sess *scs.SessionManager, engine render.Renderer) *API {
	return &API{
		store: db,
		log: ls,
		validator: vd,
		session: sess,
		engine: engine,
	}
}


func (a *API) Init() {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)	
	//	add our session middleware
	mux.Use(a.SessionLoad)

	routes := a.routes(mux)
	a.router = routes
}

func (a *API) ListenAndServe(port string) {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: a.router,
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	a.log.Info("server started", "PORT", port)
	log.Fatal(srv.ListenAndServe())
}
