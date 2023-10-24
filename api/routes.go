package api

func (a *API) AddPageRoutes() {
	a.router.Get("/", a.HandleGetHome)
}