package api

import "net/http"

func (a *API) SessionLoad(next http.Handler) http.Handler {
	a.log.Info("SessionLoad called")
	return a.session.LoadAndSave(next)
}