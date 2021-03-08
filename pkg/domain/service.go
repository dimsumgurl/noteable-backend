package domain

import "github.com/dimsumgurl/noteable-backend/pkg/api/models"

type UserService interface {
	AddUser(*models.UserAuth) error
	AuthenticateUser() error
	GetUserByID() error
	GetUserByEmail() error
	ModifyUser() error
	RemoveUser() error
}
