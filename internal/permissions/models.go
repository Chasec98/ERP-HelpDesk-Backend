package permissions

import (
	"database/sql"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PermissionSQL struct {
	ID   sql.NullInt64
	Name sql.NullString
}

func (p Permission) ToSQL() PermissionSQL {
	return PermissionSQL{
		ID:   SQLTools.ConvertInt(p.ID),
		Name: SQLTools.ConvertString(p.Name),
	}
}

func (p PermissionSQL) ToPermission() Permission {
	return Permission{
		ID:   int(p.ID.Int64),
		Name: p.Name.String,
	}
}
