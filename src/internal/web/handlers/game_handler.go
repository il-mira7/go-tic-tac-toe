package handlers

import (
	"net/http"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/web/mappers"
	webModels "tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameManager interfaces.GameManager
}

func NewGameHandler(gameManager interfaces.GameManager) *GameHandler {
	return &GameHandler{
		gameManager: gameManager,
	}
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "User not authenticated"})
		return
	}

	var req struct {
		IsVsBot bool `json:"is_vs_bot"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "Invalid request body"})
		return
	}

	game, err := h.gameManager.CreateGame(ctx, userID, req.IsVsBot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mappers.ToWebResponse(game))
}

func (h *GameHandler) GetAvailableGames(c *gin.Context) {
	ctx := c.Request.Context()

	games, err := h.gameManager.GetAvailableGames(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappers.ToWebResponseList(games))
}

func (h *GameHandler) JoinGame(c *gin.Context) {
	ctx := c.Request.Context()
	gameUUID := c.Param("uuid")
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "User not authenticated"})
		return
	}

	err := h.gameManager.JoinGame(ctx, gameUUID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined the game"})
}

func (h *GameHandler) HandleMove(c *gin.Context) {
	ctx := c.Request.Context()
	gameUUID := c.Param("uuid")
	if gameUUID == "" {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "UUID is required"})
		return
	}
	userID := c.GetString("userID")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, webModels.ErrorResponse{Message: "User not authenticated"})
		return
	}

	var req webModels.GameUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: "Invalid JSON format"})
		return
	}

	game, err := h.gameManager.MakeMove(ctx, gameUUID, userID, models.Move{Row: req.Row, Col: req.Col})
	if err != nil {
		c.JSON(http.StatusBadRequest, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappers.ToWebResponse(game))
}

func (h *GameHandler) GetGame(c *gin.Context) {
	ctx := c.Request.Context()
	gameUUID := c.Param("uuid")

	game, err := h.gameManager.GetGame(ctx, gameUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, webModels.ErrorResponse{Message: "Game not found"})
		return
	}

	c.JSON(http.StatusOK, mappers.ToWebResponse(game))
}
