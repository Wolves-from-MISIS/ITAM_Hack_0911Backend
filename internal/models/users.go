package models

type User struct {
	ID          int    `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Skills      string `json:"skills"`
	Github      string `json:"github"`
	TeamId      int    `json:"team_id"`
}

type LogInRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
