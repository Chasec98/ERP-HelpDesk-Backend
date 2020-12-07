package roles

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"

	"github.com/go-chi/chi"
)

type Api interface {
	GetRole(w http.ResponseWriter, r *http.Request)
	PostRole(w http.ResponseWriter, r *http.Request)
	GetRoles(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) GetRole(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), RoleCtxKey, Role{ID: roleID})

	role, err := a.service.GetRole(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, role)
}

func (a api) PostRole(w http.ResponseWriter, r *http.Request) {
	var role Role
	err := tools.Bind(r, &role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), RoleCtxKey, role)

	role, err = a.service.CreateRole(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, role)
}

func (a api) GetRoles(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	roles, err := a.service.GetRoles(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, roles)
}
