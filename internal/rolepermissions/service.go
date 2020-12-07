package rolepermissions

import (
	"context"
)

type Service interface {
	CreateRolesPermission(ctx context.Context) (RolesPermissions, error)
	GetRolesPermissions(ctx context.Context) ([]int, error)
	DeleteRolesPermission(ctx context.Context) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) CreateRolesPermission(ctx context.Context) (RolesPermissions, error) {
	return s.repository.CreateRolesPermission(ctx)
}

func (s service) GetRolesPermissions(ctx context.Context) ([]int, error) {
	return s.repository.GetRolesPermissions(ctx)
}

func (s service) DeleteRolesPermission(ctx context.Context) error {
	return s.repository.DeleteRolesPermission(ctx)
}
