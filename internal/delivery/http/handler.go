package http_delivery

import (
	"context"
	"github.com/gin-gonic/gin"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

type UserService interface {
	LogIn(ctx context.Context, username, mail, password string) error
	Registration(ctx context.Context, user models.RegistrationUserRequest) error
	GetRequestToFindTeam(ctx context.Context)
	RequestToApplyToTeam(ctx context.Context, teamID int)
}

type TeamService interface {
	CreateTeam(ctx context.Context)
	InviteUser(ctx context.Context)
	RemoveUser(ctx context.Context)
	TakePartInHack(ctx context.Context)
}

type Handler struct {
	userService UserService
	teamService TeamService
}

func NewHandler(userService UserService, teamService TeamService) *Handler {
	return &Handler{
		userService: userService,
		teamService: teamService,
	}
}

func (h *Handler) Init(swaggerUrl string) *gin.Engine {
	r := gin.Default()

	// api group v1
	api := r.Group("/api/v1")

	api.GET("/login", h.LogIn)
	api.GET("/registration", h.UserRegistration)

	return r
}
