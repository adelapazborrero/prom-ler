package users

import "database/sql"

type UserStore struct {
	db *sql.DB
}

func (u UserStore) Insert(user *User) error {
	return nil
}

func (u UserStore) GetById(id string) *User {
	return nil
}
