package domain

import "github.com/dimsumgurl/noteable-backend/pkg/api/models"

// UserRepo defines the interface for UserRepo
type UserRepo interface {
	CreateUserCollection() error
	CreateUser(*models.UserAuth) error
	RetrieveUserByID(int64) (*models.UserAuth, error)
	RetrieveUserByEmail(string) (*models.UserAuth, error)
	Update() error
	Delete() error
}
