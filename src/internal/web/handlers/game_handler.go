package handlers

import (
	"net/http"
	"tic-tac-toe/internal/domain/interfaces"
	domain "tic-tac-toe/internal/domain/models"
	"tic-tac-toe/internal/web/mappers"
	web "tic-tac-toe/internal/web/models"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service interfaces.GameService
}

func NewGameHandler(service interfaces.GameService) *GameHandler {
	return &GameHandler{
		service: service,
	}
}

func (h *GameHandler) HandleMove(c *gin.Context) {

	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: "UUID is required"})
		return
	}

	var req web.GameUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: "Invalid JSON format"})
		return
	}

	domainGame := mappers.ToDomain(uuid, req)
	var status domain.GameStatus
	ok, err := h.service.Validation(&domainGame, domain.Move{Row: req.Row, Col: req.Col})
	if !ok {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error()})
		return
	}

	status = h.service.CompletionCheck(domainGame)

	if status == domain.Ongoing {
		h.service.CalculateNextMove(&domainGame)
		status = h.service.CompletionCheck(domainGame)
	}

	response := mappers.ToResponse(domainGame, status)
	c.JSON(http.StatusOK, response)
}
