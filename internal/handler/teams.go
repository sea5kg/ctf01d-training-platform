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

func (h *Handler) CreateTeam(c *gin.Context) {
	var req httpserver.CreateTeamJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeam")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newTeam := &model.Team{
		Name:         req.Name,
		SocialLinks:  helper.ToNullString(req.SocialLinks),
		Description:  "",
		UniversityId: req.UniversityId,
		AvatarUrl:    helper.ToNullString(req.AvatarUrl),
	}
	if req.Description != nil {
		newTeam.Description = *req.Description
	}

	teamRepo := repository.NewTeamRepository(h.DB)
	if err := teamRepo.Create(c.Request.Context(), newTeam); err != nil {
		slog.Warn(err.Error(), "handler", "CreateTeam")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}
	c.JSON(http.StatusOK, newTeam.ToResponse())
}

func (h *Handler) DeleteTeam(c *gin.Context, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	if err := teamRepo.Delete(c.Request.Context(), id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteTeam")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete team"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Team deleted successfully"})
}

func (h *Handler) GetTeamById(c *gin.Context, id openapi_types.UUID) {
	teamRepo := repository.NewTeamRepository(h.DB)
	team, err := teamRepo.GetById(c.Request.Context(), id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetTeamById")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch team"})
		return
	}
	c.JSON(http.StatusOK, team.ToResponse())
}

func (h *Handler) ListTeams(c *gin.Context) {
	teamRepo := repository.NewTeamRepository(h.DB)
	teams, err := teamRepo.List(c.Request.Context())
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListTeams")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
		return
	}
	c.JSON(http.StatusOK, model.NewTeamsFromModels(teams))
}

func (h *Handler) UpdateTeam(c *gin.Context, id openapi_types.UUID) {
	var req httpserver.UpdateTeamJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeam")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	updateTeam := &model.Team{
		Id:           id,
		Name:         req.Name,
		SocialLinks:  helper.ToNullString(req.SocialLinks),
		Description:  "",
		UniversityId: req.UniversityId,
		AvatarUrl:    helper.ToNullString(req.AvatarUrl),
	}
	if req.Description != nil {
		updateTeam.Description = *req.Description
	}

	teamRepo := repository.NewTeamRepository(h.DB)
	if err := teamRepo.Update(c.Request.Context(), updateTeam); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateTeam")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update team"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Team updated successfully"})
}

func (h *Handler) ConnectUserTeam(c *gin.Context, teamId openapi_types.UUID, userId openapi_types.UUID) {
	var req struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	if err := teamRepo.ConnectUserTeam(c.Request.Context(), teamId, userId, req.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User create request to join the team"})
}

func (h *Handler) ApproveUserTeam(c *gin.Context, teamId openapi_types.UUID, userId openapi_types.UUID) {
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	if err := teamRepo.ApproveUserTeam(c.Request.Context(), teamId, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User approved and added to team successfully"})
}

func (h *Handler) LeaveUserFromTeam(c *gin.Context, teamId openapi_types.UUID, userId openapi_types.UUID) {
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	if err := teamRepo.LeaveUserFromTeam(c.Request.Context(), teamId, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User removed from team successfully"})
}

func (h *Handler) TeamMembers(c *gin.Context, teamId openapi_types.UUID) {
	teamRepo := repository.NewTeamMemberRequestRepository(h.DB)
	members, err := teamRepo.TeamMembers(c.Request.Context(), teamId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.NewUsersFromModels(members))
}
