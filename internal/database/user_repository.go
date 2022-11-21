package database

import (
	"fmt"

	"github.com/srselivan/user-balance-microservice/internal/model"
)

type UserRepository struct {
	db *database
}

func GetUserById(id int64) (*model.User, error) {
	var result model.User

	row := instance.store.QueryRow("SELECT * FROM users WHERE id = ?", id)

	if err := row.Scan(&result.ID, &result.Balance); err != nil {
		return nil, err
	}

	return &result, nil
}

func AppendBalanceById(id int64, amount float64) error {
	err := instance.store.Ping()
	if err != nil {
		return err
	}

	var balance float64
	row := instance.store.QueryRow("SELECT balance FROM users WHERE id = ?", id)
	if err := row.Scan(&balance); err != nil {
		return err
	}

	return instance.store.QueryRow("update users set balance = ? where id = ?", balance+amount, id).Err()
}

func TransferBalance(sendId int64, receiveId int64, amount float64) error {
	err := instance.store.Ping()
	if err != nil {
		return err
	}

	var sendIdBalance float64
	row := instance.store.QueryRow("SELECT balance FROM users WHERE id = ?", sendId)
	if err := row.Scan(&sendIdBalance); err != nil {
		return err
	}

	var receiveIdBalance float64
	row = instance.store.QueryRow("SELECT balance FROM users WHERE id = ?", receiveId)
	if err := row.Scan(&receiveIdBalance); err != nil {
		return err
	}

	if sendIdBalance < amount {
		return fmt.Errorf("Sender does not have the required amount")
	}

	err = instance.store.QueryRow("update users set balance = ? where id = ?", sendIdBalance-amount, sendId).Err()
	if err != nil {
		return err
	}

	return instance.store.QueryRow("update users set balance = ? where id = ?", receiveIdBalance+amount, receiveId).Err()
}
