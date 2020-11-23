package tickets

import (
	"database/sql"
	"time"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const TicketCtxKey = "tickets"

type Ticket struct {
	ID           int    `json:"id"`
	AssignedToID int    `json:"assignedToId"`
	CreatedByID  int    `json:"createdById"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
	Solution     string `json:"solution"`
	CreatedDate  string `json:"createdDate"`
	ClosedDate   string `json:"closedDate"`
}

type TicketSQL struct {
	ID           sql.NullInt64
	AssignedToID sql.NullInt64
	CreatedByID  sql.NullInt64
	Subject      sql.NullString
	Body         sql.NullString
	Solution     sql.NullString
	CreatedDate  sql.NullTime
	ClosedDate   sql.NullTime
}

func (t TicketSQL) ToTicket() Ticket {
	createdDate := t.CreatedDate.Time.Format(time.RFC3339)
	if t.CreatedDate.Time.IsZero() {
		createdDate = ""
	}
	closedDate := t.ClosedDate.Time.Format(time.RFC3339)
	if t.ClosedDate.Time.IsZero() {
		closedDate = ""
	}

	return Ticket{
		ID:           int(t.ID.Int64),
		AssignedToID: int(t.AssignedToID.Int64),
		CreatedByID:  int(t.CreatedByID.Int64),
		Subject:      t.Subject.String,
		Body:         t.Body.String,
		Solution:     t.Solution.String,
		CreatedDate:  createdDate,
		ClosedDate:   closedDate,
	}
}

func (t Ticket) ToTicketSQL() TicketSQL {
	var createdDateSQL = sql.NullTime{Valid: false}
	createdDate, err := time.Parse(time.RFC3339, t.CreatedDate)
	if err == nil {
		createdDateSQL = sql.NullTime{Valid: true, Time: createdDate}
	}
	var closedDateSQL = sql.NullTime{Valid: false}
	closedDate, err := time.Parse(time.RFC3339, t.ClosedDate)
	if err == nil {
		closedDateSQL = sql.NullTime{Valid: true, Time: closedDate}
	}

	return TicketSQL{
		ID:           SQLTools.ConvertInt(t.ID),
		AssignedToID: SQLTools.ConvertInt(t.AssignedToID),
		CreatedByID:  SQLTools.ConvertInt(t.CreatedByID),
		Subject:      SQLTools.ConvertString(t.Subject),
		Body:         SQLTools.ConvertString(t.Body),
		Solution:     SQLTools.ConvertString(t.Solution),
		CreatedDate:  createdDateSQL,
		ClosedDate:   closedDateSQL,
	}
}
