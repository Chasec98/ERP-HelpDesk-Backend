package auth

import (
	"database/sql"
)

const SessionCtxKey = "session"

type Session struct {
	ID        int    `sql:"ID"`
	SessionID string `sql:"SessionID"`
	UserID    int    `sql:"UserID"`
	Expires   string `sql:"Expires"`
}

type SessionSQL struct {
	ID        sql.NullInt64
	SessionID sql.NullString
	UserID    sql.NullInt64
	Expires   sql.NullTime
}

type Roles struct {
	ID          int
	Name        string
	Permissions []string
}

type RolesSQL struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Permissions []sql.NullString
}
