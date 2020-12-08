package permissions

import "context"

type Service interface {
	GetPermissions(ctx context.Context) ([]Permission, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) GetPermissions(ctx context.Context) ([]Permission, error) {
	return s.repository.GetPermissions(ctx)
}
