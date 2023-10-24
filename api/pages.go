package api

import (
	"net/http"
)

func (a *API) HandleGetHome(w http.ResponseWriter, r *http.Request)  { 
	if err := a.engine.Page(w, r, "home", nil, nil); err != nil {
		a.log.Error("error loading home page", "ERROR", err)
	}
}

func (a *API) HandleGetLogin(w http.ResponseWriter, r *http.Request) {
	if err := a.engine.Page(w, r, "login", nil, nil); err != nil {
		a.log.Error("error loading login page", "ERROR", err)
	}
}
