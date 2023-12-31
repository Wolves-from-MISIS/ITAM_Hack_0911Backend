package http_delivery

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

//	LogIn(ctx context.Context, username, mail, password string)
//	Registration(ctx context.Context, username, email, password string)
//	GetRequestToFindTeam(ctx context.Context)

func (h *Handler) LogIn(c *gin.Context) {
	var login models.LogInRequest

	session := sessions.Default(c)

	// Проверка наличия сессии
	if userID := session.Get("user_id"); userID != nil {
		// Пользователь уже авторизован, выполните действия для авторизованных пользователей.
		// Например, перенаправьте его на личную страницу или верните сообщение о том, что он уже авторизован.
		c.JSON(http.StatusOK, gin.H{"message": "Вы уже авторизованы."})
		return
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.validateLogIn(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	err = h.userService.LogIn(c, login.Username, login.Email, login.Password)
	if err != nil {
		if errors.Is(err, models.InvalidUserOrPassword) {
			c.JSON(http.StatusUnauthorized, models.ResponseMessage{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
		return
	}

	c.JSON(200, models.SucceedMessage)
	//c.Redirect(200, "/home") TODO
}

func (h *Handler) UserRegistration(c *gin.Context) {
	var login models.RegistrationUserRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.validateRegistration(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	err = h.userService.Registration(c, login)
	if err != nil {
		if errors.Is(err, models.UserAlreadyExistErr) {
			c.JSON(http.StatusSeeOther, err)
			return
		}
		c.JSON(http.StatusInternalServerError, models.ResponseMessage{Message: err.Error()})
		return
	}

	c.JSON(200, models.SucceedMessage)

	//c.Redirect(200, "/home") TODO
}

func (h *Handler) FillOutUserProfile(c *gin.Context) {
	// заполнение информации пользователя
}

func (h *Handler) UserRequireTeam(c *gin.Context) {
	// пользователь хочет запросить чтобы ему нашли команду
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

func (h *Handler) validateRegistration(l models.RegistrationUserRequest) error {
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
