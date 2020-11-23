package auth

import (
	"net/http"

	"github.com/Chasec98/ERP-HelpDesk-Backend/internal/constants"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"

	"github.com/go-chi/chi"
)

type Api interface {
	postLogin(w http.ResponseWriter, r *http.Request)
	postLogout(w http.ResponseWriter, r *http.Request)
	getSession(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func AuthRouter(service Service) http.Handler {
	api := api{
		service: service,
	}
	r := chi.NewRouter()
	r.Post("/login", api.postLogin)
	r.Post("/logout", api.postLogout)
	r.Get("/session", api.getSession)
	return r
}

func (a api) postLogin(w http.ResponseWriter, r *http.Request) {

}

func (a api) postLogout(w http.ResponseWriter, r *http.Request) {

}

func (a api) getSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(constants.CookieName)
	if err != nil {
		tools.StringReponse(w, "cookie error", http.StatusBadRequest)
		return
	}
	session, err := a.service.GetSession(Session{SessionID: cookie.Value})
	if err != nil {
		tools.StringReponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, session)
}
