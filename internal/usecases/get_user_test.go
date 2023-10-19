package usecases_test

import (
	"github.com/altheaix/RealApiWithCA/internal/infrastructure/repository"
	"github.com/altheaix/RealApiWithCA/internal/usecases"
	"testing"
)

func TestGetUserUseCase_GetUsers(t *testing.T) {
	postRepo := repository.NewUserRepository()
	createUserUseCase := usecases.NewCreateUserUseCase(postRepo)
	getUserUseCase := usecases.NewGetUserCase(postRepo)

	userInput := &usecases.CreateUserInput{Id: 1, Username: "Test", Password: "Test", Role: "Admin"}
	err := createUserUseCase.Create(userInput)
	if err != nil {
		t.Fatalf("[Error] %v", err)
	}
	userInput = &usecases.CreateUserInput{Id: 2, Username: "Test", Password: "Test", Role: "Admin"}
	_ = createUserUseCase.Create(userInput)

	userInput = &usecases.CreateUserInput{Id: 3, Username: "Test", Password: "Test", Role: "Admin"}
	_ = createUserUseCase.Create(userInput)

	users, err := getUserUseCase.GetUsers()
	if err != nil {
		t.Fatalf("[Error] %v", err)
	}

	t.Log(users)
}
