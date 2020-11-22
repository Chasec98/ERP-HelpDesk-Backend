package auth

import (
	"database/sql"
	"time"
)

type Authentication struct {
	Username string
	Password string
}

type Session struct {
	ID      int
	UserID  int
	Roles   []int
	Expires time.Time
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
