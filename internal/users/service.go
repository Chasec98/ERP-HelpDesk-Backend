package users

import (
	"context"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/pagination"
)

type Service interface {
	CreateUser(ctx context.Context) (User, error)
	UpdateUser(ctx context.Context) (User, error)
	GetUserByID(ctx context.Context) (User, error)
	GetUsers(ctx context.Context) (pagination.Pagination, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) CreateUser(ctx context.Context) (User, error) {
	return s.repository.CreateUser(ctx)
}

func (s service) UpdateUser(ctx context.Context) (User, error) {
	return s.repository.UpdateUser(ctx)
}

func (s service) GetUserByID(ctx context.Context) (User, error) {
	return s.repository.GetUserByID(ctx)
}

func (s service) GetUsers(ctx context.Context) (pagination.Pagination, error) {
	return s.repository.GetUsers(ctx)
}
