package handler

import (
	"log/slog"
	"net/http"

	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateGame(c *gin.Context) {
	var game httpserver.CreateGameJSONRequestBody
	if err := c.ShouldBindJSON(&game); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if game.EndTime.Before(game.StartTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "EndTime must be after StartTime"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	newGame := &model.Game{
		StartTime: game.StartTime,
		EndTime:   game.EndTime,
	}

	if game.Description != nil {
		newGame.Description = *game.Description
	}

	if err := repo.Create(c.Request.Context(), newGame); err != nil {
		slog.Warn(err.Error(), "handler", "CreateGame")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}
	c.JSON(http.StatusOK, newGame.ToResponse())
}

func (h *Handler) DeleteGame(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	if err := repo.Delete(c.Request.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteGame")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Game deleted successfully"})
}

func (h *Handler) GetGameById(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewGameRepository(h.DB)
	game, err := repo.GetGameDetails(c.Request.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetGameById")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch game"})
		return
	}
	c.JSON(http.StatusOK, game.ToResponseGameDetails())
}

func (h *Handler) ListGames(c *gin.Context) {
	repo := repository.NewGameRepository(h.DB)
	games, err := repo.ListGamesDetails(c.Request.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListGames")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch games"})
		return
	}
	gameResponses := make([]*httpserver.GameResponse, 0, len(games))
	for _, game := range games {
		gameResponses = append(gameResponses, game.ToResponseGameDetails())
	}
	c.JSON(http.StatusOK, gameResponses)
}

func (h *Handler) UpdateGame(c *gin.Context, id openapi_types.UUID) {
	var game httpserver.UpdateGameJSONRequestBody

	if err := c.ShouldBindJSON(&game); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewGameRepository(h.DB)
	updateGame := &model.Game{
		Id:        id,
		StartTime: game.StartTime,
		EndTime:   game.EndTime,
	}
	if game.Description != nil {
		updateGame.Description = *game.Description
	}

	if err := repo.Update(c.Request.Context(), updateGame); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateGame")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Game updated successfully"})
}
