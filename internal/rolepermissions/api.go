package rolepermissions

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
	"github.com/go-chi/chi"
)

type Api interface {
	GetRolesPermissions(w http.ResponseWriter, r *http.Request)
	PostRolesPermissions(w http.ResponseWriter, r *http.Request)
	DeleteRolesPermissions(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) GetRolesPermissions(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), RolesPermissionsCtxKey, RolesPermissions{RoleID: roleID})

	rolespermissions, err := a.service.GetRolesPermissions(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, rolespermissions)
}

func (a api) PostRolesPermissions(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	permissionID, err := strconv.Atoi(chi.URLParam(r, "permissionID"))
	if err != nil {
		http.Error(w, "Invalid Permission ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), RolesPermissionsCtxKey, RolesPermissions{RoleID: roleID, PermissionsID: permissionID})

	rolespermission, err := a.service.CreateRolesPermission(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, rolespermission)
}

func (a api) DeleteRolesPermissions(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		http.Error(w, "Invalid Role ID", http.StatusBadRequest)
		return
	}

	permissionID, err := strconv.Atoi(chi.URLParam(r, "permissionID"))
	if err != nil {
		http.Error(w, "Invalid Permission ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), RolesPermissionsCtxKey, RolesPermissions{RoleID: roleID, PermissionsID: permissionID})

	err = a.service.DeleteRolesPermission(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}
