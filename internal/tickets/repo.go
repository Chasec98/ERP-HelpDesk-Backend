package tickets

import (
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateTicket(ticket Ticket) (Ticket, error)
	UpdateTicket(ticket Ticket) (Ticket, error)
	GetTicketByID(id int) (Ticket, error)
	GetTickets() ([]Ticket, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateTicket(ticket Ticket) (Ticket, error) {
	query := `INSERT INTO Tickets (AssignedToID, CreatedByID, Subject, Body, Solution, CreatedDate, ClosedDate) VALUES (?,?,?,?,?,?,?)`
	var ticketSQL = ticket.ToSQL()
	row, err := r.db.Exec(query, ticketSQL.AssignedToID, ticketSQL.CreatedByID, ticketSQL.Subject, ticketSQL.Body, ticketSQL.Solution, ticketSQL.CreatedDate, ticketSQL.ClosedDate)
	if err != nil {
		logger.Error.Println(err.Error())
		return Ticket{}, err
	}
	ticketID, err := row.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return Ticket{}, err
	}

	return r.GetTicketByID(int(ticketID))
}

func (r repository) UpdateTicket(ticket Ticket) (Ticket, error) {
	query := `UPDATE Tickets SET AssignedToID = ?, CreatedByID = ?, Subject = ?, Body = ?, Solution = ?, CreatedDate = ?, ClosedDate = ? WHERE ID = ?`
	ticketSQL := ticket.ToSQL()
	_, err := r.db.Exec(query, ticketSQL.AssignedToID, ticketSQL.CreatedByID, ticketSQL.Subject, ticketSQL.Body, ticketSQL.Solution, ticketSQL.CreatedDate, ticketSQL.ClosedDate, ticketSQL.ID)
	if err != nil {
		logger.Error.Println(err.Error())
		return Ticket{}, err
	}
	return r.GetTicketByID(ticket.ID)
}

func (r repository) GetTicketByID(id int) (Ticket, error) {
	query := `SELECT ID, AssignedToID, CreatedByID, Subject, Body, Solution, CreatedDate, ClosedDate FROM Tickets WHERE ID = ?`
	row := r.db.QueryRow(query, id)
	var ticket = TicketSQL{}
	err := row.Scan(&ticket.ID, &ticket.AssignedToID, &ticket.CreatedByID, &ticket.Subject, &ticket.Body, &ticket.Solution, &ticket.CreatedDate, &ticket.ClosedDate)
	if err != nil {
		logger.Error.Println(err.Error())
	}
	return ticket.ToTicket(), err
}

func (r repository) GetTickets() ([]Ticket, error) {
	return []Ticket{}, nil
}
