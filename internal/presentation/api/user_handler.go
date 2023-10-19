package api

import (
	"encoding/json"
	"github.com/altheaix/RealApiWithCA/internal/domain"
	"github.com/altheaix/RealApiWithCA/internal/usecases"
	"net/http"
)

type UserHandler struct {
	CreateUserUseCase usecases.CreateUserUseCase
	GetUserUseCase    usecases.GetUserUseCase
}

func NewUserHandler(createUserUseCase usecases.CreateUserUseCase, getUserUseCase usecases.GetUserUseCase) *UserHandler {
	return &UserHandler{CreateUserUseCase: createUserUseCase, GetUserUseCase: getUserUseCase}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.CreateUserUseCase.Create(&usecases.CreateUserInput{Id: user.Id, Username: user.Username, Password: user.Password, Role: user.Role})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := u.GetUserUseCase.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
}
