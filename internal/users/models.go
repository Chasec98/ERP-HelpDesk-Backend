package users

import (
	"database/sql"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const UserCtxKey = "users"

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password,omitempty"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Active      bool   `json:"active"`
}

type UserSQL struct {
	ID          sql.NullInt64
	Username    sql.NullString
	Password    sql.NullString
	FirstName   sql.NullString
	LastName    sql.NullString
	Email       sql.NullString
	PhoneNumber sql.NullString
	Active      sql.NullBool
}

func (u User) ToUserSQL() UserSQL {
	ret := UserSQL{
		ID:          SQLTools.ConvertInt(u.ID),
		Username:    SQLTools.ConvertString(u.Username),
		Password:    SQLTools.ConvertString(u.Password),
		FirstName:   SQLTools.ConvertString(u.FirstName),
		LastName:    SQLTools.ConvertString(u.LastName),
		Email:       SQLTools.ConvertString(u.Email),
		PhoneNumber: SQLTools.ConvertString(u.PhoneNumber),
		Active:      SQLTools.ConvertBool(u.Active),
	}

	return ret
}

func (u UserSQL) ToUser() User {
	ret := User{
		ID:          int(u.ID.Int64),
		Username:    u.Username.String,
		Password:    u.Password.String,
		FirstName:   u.FirstName.String,
		LastName:    u.LastName.String,
		Email:       u.Email.String,
		PhoneNumber: u.PhoneNumber.String,
		Active:      u.Active.Bool,
	}

	return ret
}
