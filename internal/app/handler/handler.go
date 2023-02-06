package handler

import (
	"github.com/srselivan/user-balance-microservice/internal/app/service"
)

type Handler struct {
	service *service.UserBalanceService
}

func New(service *service.UserBalanceService) *Handler {
	return &Handler{
		service: service,
	}
}
