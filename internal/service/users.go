package service

import (
	"context"
	"log"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

func (u *UserService) LogIn(ctx context.Context, username, mail, password string) error {
	if username != "" {
		return u.repo.LogInByUsername(ctx, username, password)
	}

	return u.repo.LogInByMail(ctx, mail, password)
}

func (u *UserService) Registration(ctx context.Context, user models.RegistrationUserRequest) error {
	return u.repo.RegisterUser(ctx, user)
}

func (u *UserService) GetRequestToFindTeam(ctx context.Context) {
	//TODO implement me
	log.Println("implement me")
}

func (u *UserService) RequestToApplyToTeam(ctx context.Context, teamID int) {
	//TODO implement me
	log.Println("implement me")
}
