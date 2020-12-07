package rolepermissions

import (
	"database/sql"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const RolesPermissionsCtxKey = "rolespermissions"

type RolesPermissions struct {
	RoleID        int
	PermissionsID int
}

type RolesPermissionsSQL struct {
	RoleID        sql.NullInt64
	PermissionsID sql.NullInt64
}

func (r RolesPermissions) ToSQL() RolesPermissionsSQL {
	return RolesPermissionsSQL{
		RoleID:        SQLTools.ConvertInt(r.RoleID),
		PermissionsID: SQLTools.ConvertInt(r.PermissionsID),
	}
}

func (r RolesPermissionsSQL) ToRolesPermissions() RolesPermissions {
	return RolesPermissions{
		RoleID:        int(r.RoleID.Int64),
		PermissionsID: int(r.PermissionsID.Int64),
	}
}
