package handler

import (
	"log/slog"
	"net/http"

	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListUniversities(c *gin.Context, params httpserver.ListUniversitiesParams) {
	repo := repository.NewUniversityRepository(h.DB)

	var (
		universities []*model.University
		err          error
	)

	if params.Term != nil && *params.Term != "" {
		universities, err = repo.Search(c.Request.Context(), *params.Term)
	} else {
		universities, err = repo.List(c.Request.Context())
	}

	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUniversities")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch universities"})
		return
	}

	c.JSON(http.StatusOK, model.NewUniversitiesFromModels(universities))
}
