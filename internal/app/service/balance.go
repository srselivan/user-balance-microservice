package service

import (
	"errors"
	"fmt"
	"github.com/srselivan/user-balance-microservice/internal/app/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{
		repo: repo,
	}
}

func (b *BalanceService) GetBalance(ID int64) (float64, error) {
	return b.repo.GetBalance(ID)
}

func (b *BalanceService) AppendBalance(ID int64, amount float64) error {
	balance, err := b.GetBalance(ID)
	if err != nil {
		return err
	}

	newBalance := balance + amount

	return b.repo.ChangeBalance(ID, newBalance)
}

func (b *BalanceService) ReduceBalance(ID int64, amount float64) error {
	balance, err := b.GetBalance(ID)
	if err != nil {
		return err
	}

	if balance < amount {
		return errors.New(fmt.Sprintf("insufficient balance (need: %f current: %f)", amount, balance))
	}

	newBalance := balance - amount

	return b.repo.ChangeBalance(ID, newBalance)
}
