package domain

type UserService interface {
	NewService() error
	AddUser() error
	AuthenticateUser() error
	GetUserByID() error
	GetUserByEmail() error
	ModifyUser() error
	RemoveUser() error
}
