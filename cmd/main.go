package main

import (
	"github.com/altheaix/RealApiWithCA/internal/infrastructure/repository"
	"github.com/altheaix/RealApiWithCA/internal/presentation/api"
	"github.com/altheaix/RealApiWithCA/internal/usecases"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	userRepo := repository.NewUserRepository()
	createUserUseCase := usecases.NewCreateUserUseCase(userRepo)
	getUserUseCase := usecases.NewGetUserCase(userRepo)
	userHandler := api.NewUserHandler(createUserUseCase, getUserUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/user/create", userHandler.Create)
	r.HandleFunc("/user", userHandler.GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
