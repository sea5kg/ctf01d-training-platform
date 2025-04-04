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

func (h *Handler) CreateResult(c *gin.Context, gameId openapi_types.UUID) {
	var result httpserver.ResultRequest
	if err := c.ShouldBindJSON(&result); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewResultRepository(h.DB)
	newResult := &model.Result{
		GameId: gameId,
		TeamId: result.TeamId,
		Score:  result.Score,
	}

	if err := repo.Create(c.Request.Context(), newResult); err != nil {
		slog.Warn(err.Error(), "handler", "CreateResultHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create result"})
		return
	}
	c.JSON(http.StatusOK, newResult.ToResponse(0))
}

func (h *Handler) GetResult(c *gin.Context, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	result, err := repo.GetById(c.Request.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetResult")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch result"})
		return
	}
	c.JSON(http.StatusOK, result.ToResponse(0))
}

func (h *Handler) UpdateResult(c *gin.Context, gameId openapi_types.UUID, resultId openapi_types.UUID) {
	var resultRequest httpserver.ResultRequest
	if err := c.ShouldBindJSON(&resultRequest); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewResultRepository(h.DB)
	result := &model.Result{
		Id:     resultId,
		GameId: gameId,
		TeamId: resultRequest.TeamId,
		Score:  resultRequest.Score,
	}

	if err := repo.Update(c.Request.Context(), result); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateResult")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update result"})
		return
	}

	c.JSON(http.StatusOK, result.ToResponse(0))
}

func (h *Handler) GetScoreboard(c *gin.Context, gameId openapi_types.UUID) {
	repo := repository.NewResultRepository(h.DB)
	results, err := repo.List(c.Request.Context(), gameId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetScoreboard")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scoreboard"})
		return
	}

	scoreboard := model.NewScoreboardFromResults(results)
	c.JSON(http.StatusOK, scoreboard)
}
