package repository

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"log"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

func (s *HackLabRepository) RegisterUser(ctx context.Context, request models.RegistrationUserRequest) error {
	// hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err)
		return err
	}

	_, err = s.db.ExecContext(ctx, `
        INSERT INTO users (name, email, password) 
        VALUES ($1, $2, $3)`,
		request.Username, request.Email, hashedPassword)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return models.UserAlreadyExistErr
			}
		}
		log.Println("Failed to insert user:", err)
		return err
	}

	return nil
}

func (s *HackLabRepository) LogInByUsername(ctx context.Context, username, password string) error {
	var hashedPassword string

	err := s.db.QueryRowContext(ctx, "SELECT password FROM users WHERE name = $1", username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.InvalidUserOrPassword
		}
		log.Println("Failed to execute query:", err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Password doesn't match:", err)
		return models.InvalidUserOrPassword
	}

	return nil
}

func (s *HackLabRepository) LogInByMail(ctx context.Context, username, password string) error {
	//TODO implement me
	log.Println("implement me")
	return nil
}

func (s *HackLabRepository) GetUserByID(ctx context.Context, userid int) (models.User, error) {
	//TODO implement me
	log.Println("implement me")
	return models.User{}, nil
}
