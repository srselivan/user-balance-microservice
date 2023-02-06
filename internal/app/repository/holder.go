package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type HolderRepository struct {
	db *sqlx.DB
}

func NewHolderRepository(db *sqlx.DB) *HolderRepository {
	return &HolderRepository{
		db: db,
	}
}

func (h *HolderRepository) FreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error {
	query := fmt.Sprint("INSERT INTO holder VALUES($1, $2, $3, $4)")

	tx, err := h.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userID, orderID, serviceID, amount)
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

func (h *HolderRepository) UnFreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error {
	query := fmt.Sprint("DELETE FROM holder WHERE user_id = $1 AND order_id = $2 AND service_id = $3")

	tx, err := h.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userID, orderID, serviceID, amount)
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
