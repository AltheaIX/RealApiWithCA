package usecases

import (
	"github.com/altheaix/RealApiWithCA/internal/domain"
)

type GetUserInput struct {
}

type GetUserUseCase interface {
	GetUsers() (*[]domain.User, error)
}

func NewGetUserCase(userRepo domain.UserRepository) GetUserUseCase {
	return &getUserUseCase{userRepo: userRepo}
}

type getUserUseCase struct {
	userRepo domain.UserRepository
}

func (g *getUserUseCase) GetUsers() (*[]domain.User, error) {
	users, err := g.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, err
}
