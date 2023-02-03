package store

import (
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/srselivan/user-balance-microservice/internal/app/model"
)

type Store struct {
	db *sqlx.DB
}

func New() *Store {
	return &Store{
		db: &sqlx.DB{},
	}
}

func (store *Store) ConnectToDB(connStr string) error {
	var err error
	store.db, err = sqlx.Connect("pgx", connStr)
	if err != nil {
		return err
	}

	return nil
}

func (store *Store) GetUserById(id int64) (*model.User, error) {
	var user model.User

	err := store.db.Get(&user, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store *Store) AppendBalanceByUserId(id int64, amount float64) error {
	user, err := store.GetUserById(id)
	if err != nil {
		return err
	}

	tx, err := store.db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Exec(`UPDATE users SET Balance = $1 WHERE Id = $2`, user.Balance+amount, id)
	tx.Commit()

	return nil
}

func (store *Store) ReserveMoney(userID int64, orderID int64, serviceID int64, amount float64) error {
	user, err := store.GetUserById(userID)
	if err != nil {
		return err
	}

	if user.Balance < amount {
		return errors.New(fmt.Sprintf("store: insufficient balance (need: %f current: %f)", amount, user.Balance))
	}

	tx, err := store.db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Exec(`INSERT INTO reserve VALUES($1, $2, $3, $4)`, userID, orderID, serviceID, amount)
	tx.Exec(`UPDATE users SET Balance = $1 WHERE Id = $2`, user.Balance-amount, userID)
	tx.Commit()

	return nil
}

func (store *Store) MakeDeal(userID int64, orderID int64, serviceID int64, amount float64) error {
	tx, err := store.db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Exec(`DELETE FROM reserve WHERE user_id = $1 AND order_id = $2 AND service_id = $3`, userID, orderID, serviceID)
	tx.Commit()

	return nil
}
