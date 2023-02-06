package service

import "github.com/srselivan/user-balance-microservice/internal/app/repository"

type User interface {
	CreateUser() (int64, error)
	DeleteUser(ID int64) error
}

type Balance interface {
	GetBalance(ID int64) (float64, error)
	AppendBalance(ID int64, amount float64) error
	ReduceBalance(ID int64, amount float64) error
}

type Holder interface {
	FreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error
	UnFreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error
}

type UserBalanceService struct {
	User
	Balance
	Holder
}

func New(store *repository.Store) *UserBalanceService {
	return &UserBalanceService{
		User:    NewUserService(store.User),
		Balance: NewBalanceService(store.Balance),
		Holder:  NewHolderService(store.Holder),
	}
}
