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
