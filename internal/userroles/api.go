package userroles

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
	"github.com/go-chi/chi"
)

type Api interface {
	GetUserRoles(w http.ResponseWriter, r *http.Request)
	DeleteUserRole(w http.ResponseWriter, r *http.Request)
	CreateUserRole(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), UserRolesCtxKey, UserRole{UserID: userID})

	userroles, err := a.service.GetUserRoles(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, userroles)
}

func (a api) DeleteUserRole(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), UserRolesCtxKey, UserRole{UserID: userID, RoleID: roleID})

	err = a.service.DeleteUserRole(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}

func (a api) CreateUserRole(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, "Invalid UserID", http.StatusBadRequest)
		return
	}

	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), UserRolesCtxKey, UserRole{UserID: userID, RoleID: roleID})

	userrole, err := a.service.CreateUserRole(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, userrole)
}
