package rolepermissions

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateRolesPermission(ctx context.Context) (RolesPermissions, error)
	GetRolesPermissions(ctx context.Context) ([]int, error)
	DeleteRolesPermission(ctx context.Context) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) CreateRolesPermission(ctx context.Context) (RolesPermissions, error) {
	rolesPermissionsContext := ctx.Value(RolesPermissionsCtxKey).(RolesPermissions)

	query := `INSERT INTO RolesPermissions (RoleID, PermissionID) VALUES (?,?)`

	_, err := r.db.Exec(query, rolesPermissionsContext.RoleID, rolesPermissionsContext.PermissionsID)
	if err != nil {
		return RolesPermissions{}, err
	}

	return RolesPermissions{RoleID: rolesPermissionsContext.RoleID, PermissionsID: rolesPermissionsContext.PermissionsID}, nil
}

func (r repository) GetRolesPermissions(ctx context.Context) ([]int, error) {
	rolesPermissionsContext := ctx.Value(RolesPermissionsCtxKey).(RolesPermissions)

	query := `SELECT PermissionID FROM RolesPermissions WHERE RoleID = ?`
	rows, err := r.db.Query(query, rolesPermissionsContext.RoleID)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	permissionIDs := []int{}
	for rows.Next() {
		var permissionID int
		err := rows.Scan(&permissionID)
		if err != nil {
			return []int{}, err
		}
		permissionIDs = append(permissionIDs, permissionID)
	}
	return permissionIDs, nil
}

func (r repository) DeleteRolesPermission(ctx context.Context) error {
	rolesPermissionsContext := ctx.Value(RolesPermissionsCtxKey).(RolesPermissions)
	query := `DELETE FROM RolesPermissions WHERE RoleID = ? and PermissionID = ?`

	_, err := r.db.Exec(query, rolesPermissionsContext.RoleID, rolesPermissionsContext.PermissionsID)

	return err
}
