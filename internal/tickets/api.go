package tickets

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
)

type Api interface {
	postTicket(w http.ResponseWriter, r *http.Request)
	getTicket(w http.ResponseWriter, r *http.Request)
	putTicket(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func TicketRouter(service Service) http.Handler {
	api := api{
		service: service,
	}
	r := chi.NewRouter()
	r.Post("/", api.postTicket)
	r.Put("/", api.putTicket)
	r.Get("/{id}", api.getTicket)
	return r
}

func (a api) postTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	err := tools.Bind(r, &ticket)
	if err != nil {
		tools.StringReponse(w, "could not parse body")
		return
	}
	ticket, err = a.service.CreateTicket(ticket)
	if err != nil {
		tools.StringReponse(w, err.Error())
		return
	}
	tools.JSONResponse(w, ticket)
	return
}

func (a api) getTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		tools.StringReponse(w, "provide valid ticket id")
		return
	}
	ticket, err := a.service.GetTicketByID(id)
	if err != nil {
		tools.StringReponse(w, err.Error())
		return
	}
	tools.JSONResponse(w, ticket)
}

func (a api) putTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	err := tools.Bind(r, &ticket)
	if err != nil {
		tools.StringReponse(w, "could not parse body")
		return
	}
	ticket, err = a.service.UpdateTicket(ticket)
	tools.JSONResponse(w, ticket)
}
