package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (u *UserStore) CreateUser() (int64, error) {
	query := fmt.Sprint("INSERT INTO users VALUES($1) RETURNING id")
	result, err := u.db.Exec(query, 0)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func (u *UserStore) DeleteUser(ID int64) error {
	query := fmt.Sprint("DELETE FROM users WHERE id = $1")
	_, err := u.db.Exec(query, ID)
	if err != nil {
		return err
	}

	return nil
}
