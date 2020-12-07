package auth

import (
	"database/sql"
	"time"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const SessionCtxKey = "session"

type Session struct {
	ID          int
	SessionID   string
	UserID      int
	Permissions []int
	Expires     string
}

type SessionSQL struct {
	ID        sql.NullInt64
	SessionID sql.NullString
	UserID    sql.NullInt64
	Expires   sql.NullTime
}

func (s Session) ToSQL() SessionSQL {
	var expiresSQL = sql.NullTime{Valid: false}
	expires, err := time.Parse(time.RFC3339, s.Expires)
	if err == nil {
		expiresSQL = sql.NullTime{Valid: true, Time: expires}
	}

	return SessionSQL{
		ID:        SQLTools.ConvertInt(s.ID),
		SessionID: SQLTools.ConvertString(s.SessionID),
		UserID:    SQLTools.ConvertInt(s.UserID),
		Expires:   expiresSQL,
	}
}

func (s SessionSQL) ToSession() Session {
	expires := s.Expires.Time.Format(time.RFC3339)
	if s.Expires.Time.IsZero() {
		expires = ""
	}

	return Session{
		ID:        int(s.ID.Int64),
		SessionID: s.SessionID.String,
		UserID:    int(s.UserID.Int64),
		Expires:   expires,
	}
}
