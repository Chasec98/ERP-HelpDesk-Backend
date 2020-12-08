package users

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateUser(ctx context.Context) (User, error)
	UpdateUser(ctx context.Context) (User, error)
	GetUserByID(ctx context.Context) (User, error)
	GetUsers(ctx context.Context) (pagination.Pagination, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) CreateUser(ctx context.Context) (User, error) {
	userContext := ctx.Value(UserCtxKey).(User)

	query := `INSERT INTO Users (FirstName, LastName, Username, Password, Email, PhoneNumber, Active) VALUES (?,?,?,?,?,?,?)`
	var userSQL = userContext.ToUserSQL()
	row, err := r.db.Exec(query, userSQL.FirstName, userSQL.LastName, userSQL.Username, userSQL.Password, userSQL.Email, userSQL.PhoneNumber, userSQL.Active)
	if err != nil {
		logger.Error.Println(err.Error())
		return User{}, err
	}
	userID, err := row.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return User{}, err
	}

	userContext.ID = int(userID)
	ctx = context.WithValue(ctx, UserCtxKey, userContext)

	return r.GetUserByID(ctx)
}

func (r repository) UpdateUser(ctx context.Context) (User, error) {
	userContext := ctx.Value(UserCtxKey).(User)

	query := `UPDATE Users SET FirstName = ?, LastName = ?, Username = ?, Password = ?, Email = ?, PhoneNumber = ?, Active = ? WHERE ID = ?`
	userSQL := userContext.ToUserSQL()
	_, err := r.db.Exec(query, userSQL.FirstName, userSQL.LastName, userSQL.Username, userSQL.Password, userSQL.Email, userSQL.PhoneNumber, userSQL.Active, userSQL.ID)
	if err != nil {
		logger.Error.Println(err.Error())
		return User{}, err
	}

	return r.GetUserByID(ctx)
}

func (r repository) GetUserByID(ctx context.Context) (User, error) {
	userContext := ctx.Value(UserCtxKey).(User)

	query := `SELECT ID, FirstName, LastName, Username, Password, Email, PhoneNumber, Active FROM Users WHERE ID = ?`
	row := r.db.QueryRow(query, userContext.ID)
	var userSQL = UserSQL{}
	err := row.Scan(&userSQL.ID, &userSQL.FirstName, &userSQL.LastName, &userSQL.Username, &userSQL.Password, &userSQL.Email, &userSQL.PhoneNumber, &userSQL.Active)
	if err != nil {
		logger.Error.Println(err.Error())
		return User{}, nil
	}
	return userSQL.ToUser(), err
}

func (r repository) GetUsers(ctx context.Context) (pagination.Pagination, error) {
	userContext := ctx.Value(UserCtxKey).(User)
	paginationContext := ctx.Value(pagination.PaginationCtxKey).(pagination.PaginationContext)

	query := `SELECT ID, FirstName, LastName, Username, Password, Email, PhoneNumber, Active FROM Users`
	countQuery := `SELECT COUNT(*) FROM Users`
	where := " WHERE 1 = 1"
	var args []interface{}
	if userContext.Email != "" {
		where += " and Email = ?" + userContext.Email
		args = append(args, userContext.Email)
	}
	if userContext.FirstName != "" {
		where += " and FirstName = ?"
		args = append(args, userContext.FirstName)
	}
	if userContext.LastName != "" {
		where += " and LastName = ?"
		args = append(args, userContext.LastName)
	}
	if userContext.ID != 0 {
		where += " and ID = ?"
		args = append(args, userContext.ID)
	}
	if userContext.PhoneNumber != "" {
		where += " and PhoneNumber = ?"
		args = append(args, userContext.PhoneNumber)
	}
	if userContext.Username != "" {
		where += " and Username = ?"
		args = append(args, userContext.Username)
	}
	if userContext.Password != "" {
		where += " and Password = ?"
		args = append(args, userContext.Password)
	}

	var ret = pagination.Pagination{
		Offset: paginationContext.Offset,
	}

	var total int
	row := r.db.QueryRow(countQuery+where, args...)
	err := row.Scan(&total)
	if err != nil {
		logger.Error.Println(err.Error())
		return pagination.Pagination{}, err
	}
	ret.Total = total

	where += " limit ? offset ?"
	args = append(args, paginationContext.Limit)
	args = append(args, paginationContext.Offset)

	rows, err := r.db.Query(query+where, args...)
	if err != nil {
		logger.Error.Println(err.Error())
		return pagination.Pagination{}, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var userSQL UserSQL
		err := rows.Scan(&userSQL.ID, &userSQL.FirstName, &userSQL.LastName, &userSQL.Username, &userSQL.Password, &userSQL.Email, &userSQL.PhoneNumber, &userSQL.Active)
		if err != nil {
			logger.Error.Println(err.Error())
			return pagination.Pagination{}, err
		}
		users = append(users, userSQL.ToUser())
		ret.Count++
	}
	ret.Data = users

	return ret, nil
}
