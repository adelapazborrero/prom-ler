package users

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type IUserStore interface {
	GetById(id string) *User
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
	user := s.store.GetById(id)
	if user == nil {
		s.logger.Info("Failed to fetch a user", "id", id)
		return nil
	}

	return user
}
