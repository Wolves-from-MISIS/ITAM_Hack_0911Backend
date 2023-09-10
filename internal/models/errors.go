package models

import "errors"

var (
	InvalidCredentialsErr = errors.New("invalid credentials")
	InvalidUserOrPassword = errors.New("invalid username or password")
	NotAllEnvErr          = errors.New("error bad config, not all environment variables")
	UserAlreadyExistErr   = errors.New("user already exists")
)
