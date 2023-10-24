package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jimbo8702/lets_get_one/types"
	"github.com/Jimbo8702/lets_get_one/util"
	"github.com/go-chi/chi"
)

//auth routes go here
func (a *API) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := a.store.User.GetByEmail(email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	matches, err := user.PasswordMatches(password)
	if err != nil {
		w.Write([]byte("Error validating password"))
		return
	}
	if !matches {
		w.Write([]byte("Invalid password"))
		return
	}
	a.session.Put(r.Context(), "userID", user.ID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *API) Logout(w http.ResponseWriter, r *http.Request) {
	a.session.RenewToken(r.Context())
	a.session.Remove(r.Context(), "userID")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (a *API) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	u := types.User {
		FirstName: "Trevor",
		LastName:"Sawler",
		Email:"me@here.com",
		Active: 1,
		Password: "password",
	}
	id, err := a.store.User.Insert(&u)
	if err != nil {
		a.log.Error("error creating user", "error", err)
		return
	}
	fmt.Fprintf(w, "%d: %s", id, u.FirstName)
}

func (a *API) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.store.User.GetAll()
	if err != nil {
		a.log.Error("error getting all users", "error", err)
		return
	}
	for _, x := range users {
		fmt.Fprint(w, x.LastName)
	}
}

func (a *API) HandleGetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	u, err := a.store.User.GetById(id)
	if err != nil {
		a.log.Error("error getting user by id", "error", err)
		return
	}
	fmt.Fprintf(w, "%d: %s", id, u.FirstName)
}

func (a *API) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	u, err := a.store.User.GetById(id)
	if err != nil {
		a.log.Error("error getting user by id", "error", err)
		return
	}
	u.LastName = util.RandomString(10)
	err = a.store.User.Update(*u)
	if err != nil {
		a.log.Error("error updating user", "error", err)
		return
	}
	fmt.Fprintf(w, "updated last name to %s", u.LastName)
}