package user

import (
	"github.com/dimsumgurl/noteable-backend/pkg/domain"
)

type Service struct {
	UserRepo domain.UserRepo
}

func NewService(repo domain.UserRepo) (*Service, error) {
	return nil, nil
}

func (s *Service) AddUser() error {
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
