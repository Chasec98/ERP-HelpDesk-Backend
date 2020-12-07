package roles

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateRole(ctx context.Context) (Role, error)
	GetRole(ctx context.Context) (Role, error)
	GetRoles(ctx context.Context) ([]Role, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) CreateRole(ctx context.Context) (Role, error) {
	roleContext := ctx.Value(RoleCtxKey).(Role)
	roleSQL := roleContext.ToSQL()

	query := `INSERT INTO Roles (Name) VALUES (?)`
	row, err := r.db.Exec(query, roleSQL.Name)
	if err != nil {
		logger.Error.Println(err.Error())
		return Role{}, err
	}
	roleID, err := row.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return Role{}, err
	}

	roleContext.ID = int(roleID)
	ctx = context.WithValue(ctx, RoleCtxKey, roleContext)

	return r.GetRole(ctx)
}

func (r repository) GetRole(ctx context.Context) (Role, error) {
	roleContext := ctx.Value(RoleCtxKey).(Role)
	roleSQL := roleContext.ToSQL()

	query := `SELECT Name FROM Roles WHERE ID = ?`
	row := r.db.QueryRow(query, roleSQL.ID)
	err := row.Scan(&roleSQL.Name)
	if err != nil {
		logger.Error.Println(err.Error())
		return Role{}, err
	}

	return roleSQL.ToRole(), nil
}

func (r repository) GetRoles(ctx context.Context) ([]Role, error) {
	query := `SELECT ID, Name FROM Roles`
	rows, err := r.db.Query(query)
	if err != nil {
		return []Role{}, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return []Role{}, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
