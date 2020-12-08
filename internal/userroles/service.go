package userroles

import "context"

type Service interface {
	CreateUserRole(ctx context.Context) (UserRole, error)
	GetUserRoles(ctx context.Context) ([]int, error)
	DeleteUserRole(ctx context.Context) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) CreateUserRole(ctx context.Context) (UserRole, error) {
	return s.repository.CreateUserRole(ctx)
}

func (s service) GetUserRoles(ctx context.Context) ([]int, error) {
	return s.repository.GetUserRoles(ctx)
}

func (s service) DeleteUserRole(ctx context.Context) error {
	return s.repository.DeleteUserRole(ctx)
}
