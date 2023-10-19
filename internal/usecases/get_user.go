package usecases

import (
	"fmt"
	"github.com/altheaix/RealApiWithCA/internal/domain"
)

type GetUserInput struct {
}

type GetUserUseCase interface {
	GetUsers() ([]domain.User, error)
}

func NewGetUserCase(userRepo domain.UserRepository) GetUserUseCase {
	return &getUserUseCase{userRepo: userRepo}
}

type getUserUseCase struct {
	userRepo domain.UserRepository
}

func (g *getUserUseCase) GetUsers() ([]domain.User, error) {
	var users []domain.User

	for i := range domain.USER {
		users = append(users, *domain.USER[i])
	}

	fmt.Println(users)

	return users, nil
}
