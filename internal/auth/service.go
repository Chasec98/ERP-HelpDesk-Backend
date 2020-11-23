package auth

type Service interface {
	GetSession(sessionQuery Session) (Session, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) GetSession(sessionQuery Session) (Session, error) {
	return Session{}, nil
}
