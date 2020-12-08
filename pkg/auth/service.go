package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	GetSession(ctx context.Context) (Session, error)
	CreateSession(ctx context.Context) (Session, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) GetSession(ctx context.Context) (Session, error) {
	return s.repository.GetSessionByID(ctx)
}

func (s service) CreateSession(ctx context.Context) (Session, error) {
	uuid := uuid.New().String()
	time := time.Now().Add(time.Duration(time.Hour * 24)).Format(time.RFC3339)
	session := Session{SessionID: uuid, Expires: time}
	ctx = context.WithValue(ctx, SessionCtxKey, session)

	return s.repository.CreateSession(ctx)
}
