package service

import "github.com/srselivan/user-balance-microservice/internal/app/repository"

type HolderService struct {
	repo repository.Holder
}

func NewHolderService(repo repository.Holder) *HolderService {
	return &HolderService{
		repo: repo,
	}
}

func (h *HolderService) FreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error {
	return h.repo.FreezeAmount(userID, orderID, serviceID, amount)
}

func (h *HolderService) UnFreezeAmount(userID int64, orderID int64, serviceID int64, amount float64) error {
	return h.repo.UnFreezeAmount(userID, orderID, serviceID, amount)
}
