package handlers

import (
	"net/http"
	"strconv"
	"tic-tac-toe/internal/domain/interfaces"
	"tic-tac-toe/internal/web/mappers"
	webModels "tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type GameInfoHandler struct {
	gameInfo interfaces.GameInfoManager
}

func NewGameInfoHandler(gameInfo interfaces.GameInfoManager) *GameInfoHandler {
	return &GameInfoHandler{
		gameInfo: gameInfo,
	}
}

func (h *GameInfoHandler) GetUserHistory(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _ := c.Get("userID")

	games, err := h.gameInfo.GetUserHistory(ctx, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappers.ToWebResponseList(games))
}

func (h *GameInfoHandler) GetLeaderboard(c *gin.Context) {
	ctx := c.Request.Context()

	limitStr := c.DefaultQuery("n", "10")
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 10
	}

	users, err := h.gameInfo.GetLeaderboard(ctx, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, webModels.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappers.LeadersToWeb(users))
}
