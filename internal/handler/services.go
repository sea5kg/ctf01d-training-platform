package handler

import (
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/httpserver"
	"ctf01d/internal/model"
	"ctf01d/internal/repository"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateService(c *gin.Context) {
	var req httpserver.CreateServiceJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn(err.Error(), "handler", "CreateService")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	repo := repository.NewServiceRepository(h.DB)
	newService := &model.Service{
		Name:        req.Name,
		Author:      req.Author,
		LogoUrl:     helper.ToNullString(req.LogoUrl),
		Description: "",
		IsPublic:    req.IsPublic,
	}

	if req.Description != nil {
		newService.Description = *req.Description
	}

	if err := repo.Create(c.Request.Context(), newService); err != nil {
		slog.Warn(err.Error(), "handler", "CreateService")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}
	c.JSON(http.StatusOK, newService.ToResponse())
}

func (h *Handler) DeleteService(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	if err := repo.Delete(c.Request.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteService")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete service"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Service deleted successfully"})
}

func (h *Handler) GetServiceById(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewServiceRepository(h.DB)
	service, err := repo.GetById(c.Request.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetServiceById")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch service"})
		return
	}
	c.JSON(http.StatusOK, service.ToResponse())
}

func (h *Handler) ListServices(c *gin.Context) {
	repo := repository.NewServiceRepository(h.DB)
	services, err := repo.List(c.Request.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListServices")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch services"})
		return
	}
	c.JSON(http.StatusOK, model.NewServiceFromModels(services))
}

func (h *Handler) UpdateService(c *gin.Context, id openapi_types.UUID) {
	var req httpserver.UpdateServiceJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewServiceRepository(h.DB)
	update := &model.Service{
		Id:          id,
		Name:        req.Name,
		Author:      req.Author,
		LogoUrl:     helper.ToNullString(req.LogoUrl),
		Description: "",
		IsPublic:    req.IsPublic,
	}
	if req.Description != nil {
		update.Description = *req.Description
	}

	if err := repo.Update(c.Request.Context(), update); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateService")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Service updated successfully"})
}

func (h *Handler) UploadChecker(c *gin.Context, id openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *Handler) UploadService(c *gin.Context, id openapi_types.UUID) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
