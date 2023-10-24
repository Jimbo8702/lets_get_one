package api

import (
	"net/http"
)

func (a *API) HandleGetHome(w http.ResponseWriter, r *http.Request)  { 
	a.engine.Page(w, r, "home", nil, nil)
}