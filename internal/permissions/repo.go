package permissions

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	GetPermissions(ctx context.Context) ([]Permission, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) GetPermissions(ctx context.Context) ([]Permission, error) {
	query := `SELECT ID, Name FROM Permissions`
	rows, err := r.db.Query(query)
	if err != nil {
		logger.Error.Println(err.Error())
		return []Permission{}, err
	}
	defer rows.Close()

	var permissions []Permission
	for rows.Next() {
		var permission PermissionSQL
		err := rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			logger.Error.Println(err.Error())
			return []Permission{}, err
		}
		permissions = append(permissions, permission.ToPermission())
	}

	return permissions, nil
}
