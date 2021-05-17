package comments

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/tools"
	"github.com/go-chi/chi"
)

type Api interface {
	PostComment(w http.ResponseWriter, r *http.Request)
	GetComment(w http.ResponseWriter, r *http.Request)
	GetComments(w http.ResponseWriter, r *http.Request)
}

type api struct {
	service Service
}

func NewApi(service Service) Api {
	return api{service: service}
}

func (a api) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	err := tools.Bind(r, &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticketID, err := strconv.Atoi(chi.URLParam(r, "ticketid"))
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest)
		return
	}
	comment.TicketID = &ticketID

	c := context.WithValue(r.Context(), CommentCtxKey, comment)

	comment, err = a.service.CreateComment(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, comment)
}

func (a api) GetComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest)
		return
	}

	ticketID, err := strconv.Atoi(chi.URLParam(r, "ticketid"))
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), CommentCtxKey, Comment{ID: &commentID, TicketID: &ticketID})

	comment, err := a.service.GetCommentByID(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, comment)
}

func (a api) GetComments(w http.ResponseWriter, r *http.Request) {
	ticketID, err := strconv.Atoi(chi.URLParam(r, "ticketid"))
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), CommentCtxKey, Comment{TicketID: &ticketID})

	comment, err := a.service.GetCommentsByTicketID(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tools.JSONResponse(w, comment)
}
