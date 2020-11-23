package users

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
	"github.com/go-chi/chi"
)

type Api interface {
	postUser(w http.ResponseWriter, r *http.Request)
	getUser(w http.ResponseWriter, r *http.Request)
	putUser(w http.ResponseWriter, r *http.Request)
	getUsers(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func UserRouter(service Service) http.Handler {
	api := api{
		service: service,
	}
	r := chi.NewRouter()
	r.Get("/{id}", api.getUser)
	r.Put("/{id}", api.putUser)
	r.Post("/", api.postUser)
	paginatedRoutes := r.Group(nil)
	paginatedRoutes.Use(pagination.PaginationCtx)
	paginatedRoutes.Get("/", api.getUsers)
	return r
}

func (a api) getUser(w http.ResponseWriter, r *http.Request) {
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

func (a api) getUsers(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	a.service.GetUsers(c)
	fmt.Println(r.Context().Value(pagination.PaginationCtxKey).(pagination.PaginationContext))
	w.Write([]byte("ok"))
}

func (a api) postUser(w http.ResponseWriter, r *http.Request) {
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
	return
}

func (a api) putUser(w http.ResponseWriter, r *http.Request) {
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
