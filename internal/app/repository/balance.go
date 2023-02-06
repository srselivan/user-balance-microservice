package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalanceStore struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) *BalanceStore {
	return &BalanceStore{
		db: db,
	}
}

func (b *BalanceStore) GetBalance(ID int64) (float64, error) {
	query := fmt.Sprint("SELECT balance FROM users WHERE id = $1")
	balance := 0.0
	err := b.db.Get(&balance, query, ID)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (b *BalanceStore) ChangeBalance(ID int64, amount float64) error {
	query := fmt.Sprint("UPDATE users SET balance = $1 WHERE id = $2")

	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, amount, ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
