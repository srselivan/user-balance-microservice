package service

import (
	"errors"
	"fmt"
	"github.com/srselivan/user-balance-microservice/internal/app/model"
	"github.com/srselivan/user-balance-microservice/internal/app/repository"
)

type HolderService struct {
	repo repository.Holder
	*BalanceService
}

func NewHolderService(repo repository.Holder, service *BalanceService) *HolderService {
	return &HolderService{
		repo:           repo,
		BalanceService: service,
	}
}

func (h *HolderService) FreezeAmount(holderStruct model.HolderStruct) error {
	if !holderStruct.Valid() {
		return errors.New("holder service: not valid request")
	}

	balance, err := h.BalanceService.GetBalance(holderStruct.UserID)
	if err != nil {
		return err
	}

	if balance < holderStruct.Amount {
		return errors.New(fmt.Sprintf("insufficient balance (need: %f current: %f)", holderStruct.Amount, balance))
	}

	return h.repo.FreezeAmount(holderStruct)
}

func (h *HolderService) UnFreezeAmount(holderStruct model.HolderStruct) error {
	if !holderStruct.Valid() {
		return errors.New("holder service: not valid request")
	}

	return h.repo.UnFreezeAmount(holderStruct)
}
