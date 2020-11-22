package tickets

import (
	"database/sql"
	"time"

	customSQL "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

type Ticket struct {
	ID           int       `json:"id"`
	AssignedToID int       `json:"assignedToId"`
	CreatedByID  int       `json:"createdById"`
	Subject      string    `json:"subject"`
	Body         string    `json:"body"`
	Solution     string    `json:"solution"`
	CreatedDate  time.Time `json:"createdDate"`
	ClosedDate   time.Time `json:"closedDate"`
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
	return Ticket{
		ID:           int(t.ID.Int64),
		AssignedToID: int(t.AssignedToID.Int64),
		CreatedByID:  int(t.CreatedByID.Int64),
		Subject:      t.Subject.String,
		Body:         t.Body.String,
		Solution:     t.Solution.String,
		CreatedDate:  t.CreatedDate.Time,
		ClosedDate:   t.ClosedDate.Time,
	}
}

func (t Ticket) ToSQL() TicketSQL {
	return TicketSQL{
		ID:           customSQL.ConvertInt(t.ID),
		AssignedToID: customSQL.ConvertInt(t.AssignedToID),
		CreatedByID:  customSQL.ConvertInt(t.CreatedByID),
		Subject:      customSQL.ConvertString(t.Subject),
		Body:         customSQL.ConvertString(t.Body),
		Solution:     customSQL.ConvertString(t.Solution),
		CreatedDate:  customSQL.ConvertTime(t.CreatedDate),
		ClosedDate:   customSQL.ConvertTime(t.ClosedDate),
	}
}
