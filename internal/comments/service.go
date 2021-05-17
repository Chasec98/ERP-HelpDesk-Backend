package comments

import "context"

type Service interface {
	CreateComment(ctx context.Context) (Comment, error)
	UpdateComment(ctx context.Context) (Comment, error)
	GetCommentByID(ctx context.Context) (Comment, error)
	GetCommentsByTicketID(ctx context.Context) ([]Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (s service) CreateComment(ctx context.Context) (Comment, error) {
	//TODO: check make sure ticket and user exists
	return s.repository.CreateComment(ctx)
}

func (s service) UpdateComment(ctx context.Context) (Comment, error) {
	return s.repository.UpdateComment(ctx)
}

func (s service) GetCommentByID(ctx context.Context) (Comment, error) {
	return s.repository.GetCommentByID(ctx)
}

func (s service) GetCommentsByTicketID(ctx context.Context) ([]Comment, error) {
	return s.repository.GetCommentsByTicketID(ctx)
}
