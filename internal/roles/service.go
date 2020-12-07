package roles

import (
	"context"
)

type Service interface {
	CreateRole(ctx context.Context) (Role, error)
	GetRole(ctx context.Context) (Role, error)
	GetRoles(ctx context.Context) ([]Role, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) CreateRole(ctx context.Context) (Role, error) {
	return s.repository.CreateRole(ctx)
}

func (s service) GetRole(ctx context.Context) (Role, error) {
	return s.repository.GetRole(ctx)
}

func (s service) GetRoles(ctx context.Context) ([]Role, error) {
	return s.repository.GetRoles(ctx)
}
