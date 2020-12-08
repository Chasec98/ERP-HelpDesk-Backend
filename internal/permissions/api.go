package permissions

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
)

type Api interface {
	GetPermissions(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) GetPermissions(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	permissions, err := a.service.GetPermissions(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, permissions)
}
