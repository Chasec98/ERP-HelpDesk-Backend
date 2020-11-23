package auth

import (
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateSession(userID int) (Session, error)
	GetSessionByID(sessionID int) (Session, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) GetSessionByID(sessionID int) (Session, error) {
	query := `SELECT ID, SessionID, UserID, Expires FROM Sessions WHERE SessionID = ?`
	row := r.db.QueryRow(query, sessionID)
	var session Session
	err := row.Scan(&session.ID, &session.SessionID, &session.UserID, &session.Expires)
	return session, err
}

func (r repository) CreateSession(userID int) (Session, error) {
	query := `INSERT INTO Sessions (SessionID, UserID, Expires) VALUES (?,?,?)`
	row, err := r.db.Exec(query)
	if err != nil {
		logger.Error.Println(err.Error())
		return Session{}, err
	}
	sessionID, err := row.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return Session{}, err
	}
	return r.GetSessionByID(int(sessionID))
}
