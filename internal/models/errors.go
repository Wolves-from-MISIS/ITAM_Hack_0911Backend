package models

import "errors"

var (
	InvalidCredentialsErr = errors.New("invalid credentials")
	InvalidUserOrPassword = errors.New("invalid username or password")
)
