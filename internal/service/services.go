package service

import (
	"context"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

type HackUserRepository interface {
	RegisterUser(ctx context.Context, request models.RegistrationUserRequest) error
	GetUserByID(ctx context.Context, userid int) (models.User, error)
	LogInByUsername(ctx context.Context, username, password string) error
	LogInByMail(ctx context.Context, username, password string) error
}

type UserService struct {
	repo HackUserRepository
}

func NewUserService(repo HackUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
