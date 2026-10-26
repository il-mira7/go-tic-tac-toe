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

	var webReq webModels.JwtRequest

	if err := c.ShouldBindJSON(&webReq); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "Invalid request body: login and password are required"})
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

	c.JSON(http.StatusCreated, webModels.RegisterResponse{Message: "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req webModels.JwtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "Invalid request body: login and password are required"})
		return
	}

	tokens, err := h.authService.Login(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, webModels.JwtResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    tokens.TokenType,
	})
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "Unauthorized"})
		return
	}

	user, err := h.userService.GetByID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, webModels.ErrorResponse{Message: "User not found"})
		return
	}

	response := mappers.ToWebUserInfoResponse(user)
	c.JSON(http.StatusOK, response)

}
func (h *AuthHandler) RefreshAccessToken(c *gin.Context) {
	var req webModels.RefreshJwtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "refresh token is required"})
		return
	}

	tokens, err := h.authService.RefreshAccessToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, webModels.JwtResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req webModels.RefreshJwtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "refresh token is required"})
		return
	}

	tokens, err := h.authService.RefreshRefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, webModels.JwtResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		TokenType:    "Bearer",
	})
}
