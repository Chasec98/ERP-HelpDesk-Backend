package auth

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateSession(ctx context.Context) (Session, error)
	GetSessionByID(ctx context.Context) (Session, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetSessionByID(ctx context.Context) (Session, error) {
	sessionContext := ctx.Value(SessionCtxKey).(Session)

	sessionSQL := sessionContext.ToSQL()

	query := `SELECT ID, SessionID, UserID, Expires FROM Sessions WHERE SessionID = ?`
	row := r.db.QueryRow(query, sessionSQL.SessionID)
	var session Session
	err := row.Scan(&session.ID, &session.SessionID, &session.UserID, &session.Expires)
	return session, err
}

func (r repository) CreateSession(ctx context.Context) (Session, error) {
	sessionContext := ctx.Value(SessionCtxKey).(Session)

	sessionSQL := sessionContext.ToSQL()

	query := `INSERT INTO Sessions (SessionID, UserID, Expires) VALUES (?,?,?)`
	_, err := r.db.Exec(query, sessionSQL.SessionID, sessionSQL.UserID, sessionSQL.Expires, sessionSQL.Expires)
	if err != nil {
		logger.Error.Println(err.Error())
		return Session{}, err
	}

	return r.GetSessionByID(ctx)
}
