package usecases_test

import (
	"github.com/altheaix/RealApiWithCA/internal/infrastructure/repository"
	"github.com/altheaix/RealApiWithCA/internal/usecases"
	"testing"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	postRepo := repository.NewUserRepository()
	userUseCase := usecases.NewCreateUserUseCase(postRepo)

	userInput := &usecases.CreateUserInput{Username: "Test", Password: "Test", Role: "Admin"}

	err := userUseCase.Create(userInput)
	if err != nil {
		t.Fatalf("[Error] %v", err)
	}
	t.Log(err)
}
