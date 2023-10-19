package repository

import "github.com/altheaix/RealApiWithCA/internal/domain"

type UserRepo struct {
}

func NewUserRepository() domain.UserRepository {
	return &UserRepo{}
}

func (u *UserRepo) Create(user *domain.User) error {
	domain.USER[user.Id] = user
	return nil
}

func (u *UserRepo) GetUsers() (*[]domain.User, error) {
	var users []domain.User

	for i := range domain.USER {
		users = append(users, *domain.USER[i])
	}

	return &users, nil
}
