package userroles

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateUserRole(ctx context.Context) (UserRole, error)
	GetUserRoles(ctx context.Context) ([]int, error)
	DeleteUserRole(ctx context.Context) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) CreateUserRole(ctx context.Context) (UserRole, error) {
	userRolesContext := ctx.Value(UserRolesCtxKey).(UserRole)
	userRolesSQL := userRolesContext.ToSQL()

	query := `INSERT INTO UserRoles (UserID, RoleID) VALUE (?,?)`

	_, err := r.db.Exec(query, userRolesSQL.UserID, userRolesSQL.RoleID)
	if err != nil {
		return UserRole{}, err
	}

	return UserRole{UserID: userRolesContext.UserID, RoleID: userRolesContext.RoleID}, nil
}

func (r repository) GetUserRoles(ctx context.Context) ([]int, error) {
	userRolesContext := ctx.Value(UserRolesCtxKey).(UserRole)
	userRolesSQL := userRolesContext.ToSQL()

	query := `SELECT RoleID FROM UserRoles WHERE UserID = ?`
	rows, err := r.db.Query(query, userRolesSQL.UserID)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	roleIDs := []int{}
	for rows.Next() {
		var roleIDSQL sql.NullInt64
		err := rows.Scan(&roleIDSQL)
		if err != nil {
			return []int{}, err
		}
		roleIDs = append(roleIDs, int(roleIDSQL.Int64))
	}

	return roleIDs, nil
}

func (r repository) DeleteUserRole(ctx context.Context) error {
	userRolesContext := ctx.Value(UserRolesCtxKey).(UserRole)
	userRolesSQL := userRolesContext.ToSQL()

	query := `DELETE FROM UserRoles WHERE UserID = ? and RoleID = ?`
	_, err := r.db.Exec(query, userRolesSQL.UserID, userRolesSQL.RoleID)

	return err
}
