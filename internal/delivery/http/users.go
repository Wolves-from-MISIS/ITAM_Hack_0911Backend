package http

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

//	LogIn(ctx context.Context, username, mail, password string)
//	Registration(ctx context.Context, username, email, password string)
//	GetRequestToFindTeam(ctx context.Context)

func (h *Handler) LogIn(c *gin.Context) {
	// getting json from request

	var login models.LogInRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.validateLogIn(login)
	if err != nil {
		c.JSON(400, gin.H{"error: ": err.Error()})
		return
	}

}

func (h *Handler) validateLogIn(l models.LogInRequest) error {
	if l.Username == "" && l.Email == "" {
		return models.InvalidCredentialsErr
	}

	if l.Username != "" {
		usernamePattern := "^[a-zA-Z0-9]+$"
		if ok, _ := regexp.MatchString(usernamePattern, l.Username); !ok {
			return models.InvalidCredentialsErr
		}
	}

	if l.Username == "" {
		if !strings.Contains(l.Email, "@") {
			return models.InvalidCredentialsErr
		}
	}

	if l.Password == "" {
		return models.InvalidCredentialsErr
	}

	return nil
}
