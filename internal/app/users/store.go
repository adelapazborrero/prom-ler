package users

import (
	"database/sql"
	"errors"
)

type UserStore struct {
	db *sql.DB
}

func (u UserStore) Insert(user *User) error {
	query := "insert into users (username, email) values ($1, $2)"

	_, err := u.db.Exec(query, user.Username, user.Email)

	if err != nil {
		return err
	}

	return nil
}

func (u UserStore) GetById(id string) (*User, error) {
	var user User

	query := "select * from users where id = $1"

	row := u.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
