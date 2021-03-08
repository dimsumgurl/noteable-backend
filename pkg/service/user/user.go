package user

import (
	"errors"

	"github.com/dimsumgurl/noteable-backend/pkg/api/models"
	"github.com/dimsumgurl/noteable-backend/pkg/domain"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepo domain.UserRepo
}

func NewService(userRepo domain.UserRepo) (*Service, error) {
	if userRepo == nil {
		return nil, errors.New("userRepo is undefined")
	}
	return &Service{UserRepo: userRepo}, nil
}

func hashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hash), nil
}
func (s *Service) AddUser(user *models.UserAuth) error {
	var err error
	user.PasswordHash, err = hashAndSalt([]byte(*user.Password))
	if err != nil {
		return err
	}
	// not storing actual passwords
	user.Password = nil

	err = s.UserRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AuthenticateUser() error {
	return nil
}

func (s *Service) GetUserByID() error {
	return nil
}

func (s *Service) GetUserByEmail() error {
	return nil
}

func (s *Service) ModifyUser() error {
	return nil
}

func (s *Service) RemoveUser() error {
	return nil
}
