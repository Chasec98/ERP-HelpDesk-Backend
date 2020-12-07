package tickets

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
)

type Api interface {
	PostTicket(w http.ResponseWriter, r *http.Request)
	GetTicket(w http.ResponseWriter, r *http.Request)
	PutTicket(w http.ResponseWriter, r *http.Request)
	GetTickets(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) PostTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	err := tools.Bind(r, &ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), TicketCtxKey, ticket)

	ticket, err = a.service.CreateTicket(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, ticket)
	return
}

func (a api) GetTicket(w http.ResponseWriter, r *http.Request) {
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

func (a api) PutTicket(w http.ResponseWriter, r *http.Request) {
	ticketID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Ticket ID", http.StatusBadRequest)
		return
	}

	var ticket Ticket
	err = tools.Bind(r, &ticket)
	if err != nil {
		http.Error(w, "could not parse body", http.StatusBadRequest)
		return
	}
	ticket.ID = ticketID

	c := context.WithValue(r.Context(), TicketCtxKey, ticket)

	ticket, err = a.service.UpdateTicket(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tools.JSONResponse(w, ticket)
}

func (a api) GetTickets(w http.ResponseWriter, r *http.Request) {

}
