package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (a *API) routes(router *chi.Mux) *chi.Mux  {
	fileServer := http.FileServer(http.Dir("./public"))
	router.Handle("/public/*", http.StripPrefix("/public", fileServer))
	router.Get("/", a.HandleGetHome)
	router.Mount("/pages", a.pageRoutes())
	router.Mount("/api", a.apiRoutes())
	return router
}

// /pages
func (a *API) pageRoutes() *chi.Mux {
	pages := chi.NewRouter()
	pages.Get("/users/login", a.HandleGetLogin)
	return pages
}

// /api
func (a *API) apiRoutes() *chi.Mux {
	api := chi.NewRouter()
	api.Post("/users/login", a.HandleLoginUser)
	api.Get("/test-create-user", a.HandleCreateUser)
	api.Get("/test-get-all-users", a.HandleGetAllUsers)
	api.Get("/test-by-id/{id}", a.HandleGetById)
	api.Get("/test-update-user/{id}", a.HandleUpdateUser)
	return api
}