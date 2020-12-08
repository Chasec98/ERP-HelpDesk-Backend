package roles

import (
	"database/sql"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const RoleCtxKey = "roles"

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoleSQL struct {
	ID   sql.NullInt64
	Name sql.NullString
}

func (r Role) ToSQL() RoleSQL {
	return RoleSQL{
		ID:   SQLTools.ConvertInt(r.ID),
		Name: SQLTools.ConvertString(r.Name),
	}
}

func (r RoleSQL) ToRole() Role {
	return Role{
		ID:   int(r.ID.Int64),
		Name: r.Name.String,
	}
}
