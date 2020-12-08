package userroles

import (
	"database/sql"

	SQLTools "github.com/Chasec98/ERP-HelpDesk-Backend/pkg/sql"
)

const UserRolesCtxKey = "userroles"

type UserRole struct {
	UserID int
	RoleID int
}

type UserRoleSQL struct {
	UserID sql.NullInt64
	RoleID sql.NullInt64
}

func (u UserRole) ToSQL() UserRoleSQL {
	return UserRoleSQL{
		UserID: SQLTools.ConvertInt(u.UserID),
		RoleID: SQLTools.ConvertInt(u.RoleID),
	}
}

func (u UserRoleSQL) ToUserRole() UserRole {
	return UserRole{
		UserID: int(u.UserID.Int64),
		RoleID: int(u.RoleID.Int64),
	}
}
