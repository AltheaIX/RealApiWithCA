package usecases

import (
	"github.com/altheaix/RealApiWithCA/internal/domain"
)

type CreateUserInput struct {
	Id       int
	Username string
	Password string
	Role     string
}

type CreateUserUseCase interface {
	Create(input *CreateUserInput) error
}

func NewCreateUserUseCase(repo domain.UserRepository) CreateUserUseCase {
	return &createUserUseCase{userRepo: repo}
}

type createUserUseCase struct {
	userRepo domain.UserRepository
}

func (c *createUserUseCase) Create(input *CreateUserInput) error {
	err := c.userRepo.Create(&domain.User{Id: input.Id, Username: input.Username, Password: input.Password, Role: input.Role})
	if err != nil {
		return err
	}

	return nil
}
