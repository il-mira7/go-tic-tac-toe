package middleware

import (
	"net/http"
	"strings"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type UserAuthenticator struct {
	tokenProvider interfaces.TokenProvider
}

func NewUserAuthenticator(tokenProv interfaces.TokenProvider) *UserAuthenticator {
	return &UserAuthenticator{
		tokenProvider: tokenProv,
	}
}

func (a *UserAuthenticator) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Message: "Authorization header is required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Message: "Authorization header must be Bearer {token}",
			})
			return
		}

		accessToken := parts[1]

		userUUID, err := a.tokenProvider.ExtractUUID(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Message: "Invalid or expired token",
			})
			return
		}

		c.Set("userID", userUUID.String())

		c.Next()
	}
}
