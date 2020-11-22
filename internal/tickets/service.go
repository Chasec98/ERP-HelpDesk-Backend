package tickets

type Service interface {
	CreateTicket(ticket Ticket) (Ticket, error)
	UpdateTicket(ticket Ticket) (Ticket, error)
	GetTicketByID(id int) (Ticket, error)
	GetTickets() ([]Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) CreateTicket(ticket Ticket) (Ticket, error) {
	return s.repository.CreateTicket(ticket)
}

func (s service) UpdateTicket(ticket Ticket) (Ticket, error) {
	return s.repository.UpdateTicket(ticket)
}

func (s service) GetTicketByID(id int) (Ticket, error) {
	return s.repository.GetTicketByID(id)
}

func (s service) GetTickets() ([]Ticket, error) {
	return s.repository.GetTickets()
}
