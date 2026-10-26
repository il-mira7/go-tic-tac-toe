package handlers

import (
	"net/http"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/web/mappers"
	webModels "tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService interfaces.AuthService
	userService interfaces.UserService
}

func NewAuthHandler(authSvc interfaces.AuthService, userSvc interfaces.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authSvc,
		userService: userSvc,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {

	var webReq webModels.SignUpRequest

	if err := c.ShouldBindJSON(&webReq); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{
			Message: "Invalid request body: login and password are required",
		})
		return
	}

	domainReq := mappers.ToDomainSignUp(webReq)

	err := h.authService.SignUp(c.Request.Context(), domainReq)
	if err != nil {
		c.JSON(http.StatusConflict, webModels.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, webModels.RegisterResponse{
		Message: "User registered successfully",
	})
}

func (h *AuthHandler) Authenticate(c *gin.Context) {
	login, password, ok := c.Request.BasicAuth()
	if !ok {
		c.Header("WWW-Authenticate", `Basic realm="restricted"`)
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{
			Message: "Authorization header is missing or malformed",
		})
		return
	}

	userUUID, err := h.authService.Authenticate(c.Request.Context(), login, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{
			Message: "Invalid login or password",
		})
		return
	}

	c.JSON(http.StatusOK, webModels.AuthResponse{
		UUID: userUUID,
	})
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	ctx := c.Request.Context()
	userUUID := c.Param("uuid")

	if userUUID == "" {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "User UUID is required"})
		return
	}

	user, err := h.userService.GetByID(ctx, userUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, webModels.ErrorResponse{Message: "User not found"})
		return
	}

	response := mappers.ToWebUserInfoResponse(user)
	c.JSON(http.StatusOK, response)
}
