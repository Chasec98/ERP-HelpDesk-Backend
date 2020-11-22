package users

type Service interface {
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	GetUserByID(id int) (User, error)
	GetUsers() ([]User, error)
}
