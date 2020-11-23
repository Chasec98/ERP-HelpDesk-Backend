package tickets

import "context"

type Service interface {
	CreateTicket(ctx context.Context) (Ticket, error)
	UpdateTicket(ctx context.Context) (Ticket, error)
	GetTicketByID(ctx context.Context) (Ticket, error)
	GetTickets(ctx context.Context) ([]Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) CreateTicket(ctx context.Context) (Ticket, error) {
	return s.repository.CreateTicket(ctx)
}

func (s service) UpdateTicket(ctx context.Context) (Ticket, error) {
	return s.repository.UpdateTicket(ctx)
}

func (s service) GetTicketByID(ctx context.Context) (Ticket, error) {
	return s.repository.GetTicketByID(ctx)
}

func (s service) GetTickets(ctx context.Context) ([]Ticket, error) {
	return s.repository.GetTickets(ctx)
}
