package middleware

import (
	"net/http"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type UserAuthenticator struct {
	authService interfaces.AuthService
}

func NewUserAuthenticator(service interfaces.AuthService) *UserAuthenticator {
	return &UserAuthenticator{
		authService: service,
	}
}

func (a *UserAuthenticator) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		login, password, ok := c.Request.BasicAuth()
		if !ok {
			c.Header("WWW-Authenticate", `Basic realm="restricted"`)
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Message: "Authorization required",
			})
			return
		}

		userId, err := a.authService.Authenticate(c.Request.Context(), login, password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Message: "Invalid login or password",
			})
			return
		}

		c.Set("user_id", userId)

		c.Next()
	}
}
