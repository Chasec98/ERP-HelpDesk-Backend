package tickets

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"

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

	paginated := r.Group(nil)
	paginated.Use(pagination.PaginationCtx)
	paginated.Get("/", api.getTickets)

	return r
}

func (a api) postTicket(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var ticket Ticket
	err := tools.Bind(r, &ticket)
	if err != nil {
		http.Error(w, "could not parse body", http.StatusBadRequest)
		return
	}
	ticket, err = a.service.CreateTicket(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, ticket)
	return
}

func (a api) getTicket(w http.ResponseWriter, r *http.Request) {
	ticketID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Ticket ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), TicketCtxKey, Ticket{ID: ticketID})

	ticket, err := a.service.GetTicketByID(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, ticket)
}

func (a api) putTicket(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var ticket Ticket
	err := tools.Bind(r, &ticket)
	if err != nil {
		http.Error(w, "could not parse body", http.StatusBadRequest)
		return
	}
	ticket, err = a.service.UpdateTicket(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, ticket)
}

func (a api) getTickets(w http.ResponseWriter, r *http.Request) {

}
