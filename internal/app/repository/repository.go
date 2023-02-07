package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/srselivan/user-balance-microservice/internal/app/model"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User interface {
	CreateUser() (int64, error)
	DeleteUser(ID int64) error
}

type Balance interface {
	GetBalance(ID int64) (float64, error)
	ChangeBalance(ID int64, amount float64) error
}

type Holder interface {
	FreezeAmount(holderStruct model.HolderStruct) error
	UnFreezeAmount(holderStruct model.HolderStruct) error
}

type Store struct {
	User
	Balance
	Holder
}

func New(db *sqlx.DB) *Store {
	return &Store{
		User:    NewUserRepository(db),
		Balance: NewBalanceRepository(db),
		Holder:  NewHolderRepository(db),
	}
}
