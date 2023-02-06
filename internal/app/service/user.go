package service

import "github.com/srselivan/user-balance-microservice/internal/app/repository"

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser() (int64, error) {
	return u.repo.CreateUser()
}

func (u *UserService) DeleteUser(ID int64) error {
	return u.repo.DeleteUser(ID)
}
