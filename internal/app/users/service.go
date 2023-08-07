package users

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type IUserStore interface {
	GetById(id string) (*User, error)
	Insert(user *User) error
}

type Service struct {
	store  IUserStore
	logger logrus.Logger
}

func NewService(db *sql.DB) *Service {
	return &Service{
		store:  UserStore{db},
		logger: *logrus.New(),
	}
}

func (s *Service) GetById(id string) *User {
	user, err := s.store.GetById(id)
	if err != nil {
		s.logger.Info("Failed to fetch a user", "id", id, "error", err)
		return nil
	}

	return user
}

func (s *Service) Insert(user *User) error {
	err := s.store.Insert(user)
	if err != nil {
		s.logger.Info("Failed to insert user", "username", user.Username, "email", user.Email)
		return nil
	}
	return nil
}
