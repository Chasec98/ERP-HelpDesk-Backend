package users

import (
	"database/sql"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Role        int
	Username    string
}

type UserSQL struct {
	ID          sql.NullInt64
	FirstName   sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	PhoneNumber sql.NullString
	Role        sql.NullString
	Username    sql.NullString
}
