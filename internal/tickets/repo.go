package tickets

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateTicket(ctx context.Context) (Ticket, error)
	UpdateTicket(ctx context.Context) (Ticket, error)
	GetTicketByID(ctx context.Context) (Ticket, error)
	GetTickets(ctx context.Context) ([]Ticket, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateTicket(ctx context.Context) (Ticket, error) {
	ticketContext := ctx.Value(TicketCtxKey).(Ticket)

	query := `INSERT INTO Tickets (AssignedToID, CreatedByID, Subject, Body, Solution, CreatedDate, ClosedDate) VALUES (?,?,?,?,?,?,?)`
	var ticketSQL = ticketContext.ToTicketSQL()
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

	ticketContext.ID = int(ticketID)
	ctx = context.WithValue(ctx, TicketCtxKey, ticketContext)

	return r.GetTicketByID(ctx)
}

func (r repository) UpdateTicket(ctx context.Context) (Ticket, error) {
	ticketContext := ctx.Value(TicketCtxKey).(Ticket)

	query := `UPDATE Tickets SET AssignedToID = ?, CreatedByID = ?, Subject = ?, Body = ?, Solution = ?, CreatedDate = ?, ClosedDate = ? WHERE ID = ?`
	ticketSQL := ticketContext.ToTicketSQL()
	_, err := r.db.Exec(query, ticketSQL.AssignedToID, ticketSQL.CreatedByID, ticketSQL.Subject, ticketSQL.Body, ticketSQL.Solution, ticketSQL.CreatedDate, ticketSQL.ClosedDate, ticketSQL.ID)
	if err != nil {
		logger.Error.Println(err.Error())
		return Ticket{}, err
	}
	return r.GetTicketByID(ctx)
}

func (r repository) GetTicketByID(ctx context.Context) (Ticket, error) {
	ticketContext := ctx.Value(TicketCtxKey).(Ticket)

	query := `SELECT ID, AssignedToID, CreatedByID, Subject, Body, Solution, CreatedDate, ClosedDate FROM Tickets WHERE ID = ?`
	row := r.db.QueryRow(query, ticketContext.ID)
	var ticketSQL = TicketSQL{}
	err := row.Scan(&ticketSQL.ID, &ticketSQL.AssignedToID, &ticketSQL.CreatedByID, &ticketSQL.Subject, &ticketSQL.Body, &ticketSQL.Solution, &ticketSQL.CreatedDate, &ticketSQL.ClosedDate)
	if err != nil {
		logger.Error.Println(err.Error())
		return Ticket{}, err
	}
	return ticketSQL.ToTicket(), err
}

func (r repository) GetTickets(ctx context.Context) ([]Ticket, error) {
	return []Ticket{}, nil
}
