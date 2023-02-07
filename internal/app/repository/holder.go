package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/srselivan/user-balance-microservice/internal/app/model"
)

type HolderRepository struct {
	db *sqlx.DB
}

func NewHolderRepository(db *sqlx.DB) *HolderRepository {
	return &HolderRepository{
		db: db,
	}
}

func (h *HolderRepository) FreezeAmount(holderStruct model.HolderStruct) error {
	holderQuery := fmt.Sprint("INSERT INTO holder VALUES($1, $2, $3, $4)")
	balanceQuery := fmt.Sprint("UPDATE users SET balance = (balance - $1) WHERE id = $2")

	tx, err := h.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(holderQuery, holderStruct.UserID, holderStruct.OrderID, holderStruct.ServiceID, holderStruct.Amount)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(balanceQuery, holderStruct.Amount, holderStruct.UserID)
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

func (h *HolderRepository) UnFreezeAmount(holderStruct model.HolderStruct) error {
	query := fmt.Sprint("DELETE FROM holder WHERE user_id = $1 AND order_id = $2 AND service_id = $3")

	tx, err := h.db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, holderStruct.UserID, holderStruct.OrderID, holderStruct.ServiceID, holderStruct.Amount)
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
