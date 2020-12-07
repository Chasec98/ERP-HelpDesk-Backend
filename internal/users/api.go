package users

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
	"github.com/go-chi/chi"
)

type Api interface {
	PostUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), UserCtxKey, User{ID: userID})

	user, err := a.service.GetUserByID(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, user)
}

func (a api) GetUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	tools.Bind(r, &user)

	c := context.WithValue(r.Context(), UserCtxKey, user)

	pagination, err := a.service.GetUsers(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, pagination)
}

func (a api) PostUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := tools.Bind(r, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), UserCtxKey, user)

	user, err = a.service.CreateUser(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, user)
}

func (a api) PutUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	var user User
	err = tools.Bind(r, &user)
	if err != nil {
		http.Error(w, "could not parse body", http.StatusBadRequest)
		return
	}
	user.ID = userID

	c := context.WithValue(r.Context(), UserCtxKey, user)

	user, err = a.service.UpdateUser(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, user)
}
