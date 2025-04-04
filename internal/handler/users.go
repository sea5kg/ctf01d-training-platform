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

func (h *Handler) CreateUser(c *gin.Context) {
	var user httpserver.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	repo := repository.NewUserRepository(h.DB)
	passwordHash, err := helper.HashPassword(user.Password)
	if err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password hash"})
		return
	}

	newUser := &model.User{
		Username:     user.UserName,
		DisplayName:  helper.ToNullString(user.DisplayName),
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    helper.ToNullString(user.AvatarUrl),
	}

	if err = repo.Create(c, newUser); err != nil {
		slog.Warn(err.Error(), "handler", "CreateUserHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	if user.TeamIds != nil && len(*user.TeamIds) > 0 {
		if err := repo.AddUserToTeams(c, newUser.Id, user.TeamIds); err != nil {
			slog.Warn(err.Error(), "handler", "CreateUserHandler")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to teams"})
			return
		}
	}

	c.JSON(http.StatusOK, newUser.ToResponse())
}

func (h *Handler) DeleteUser(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	if err := repo.Delete(c, id); err != nil {
		slog.Warn(err.Error(), "handler", "DeleteUserHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}

func (h *Handler) GetUserById(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	user, err := repo.GetById(c, id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "GetUserByIdHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	c.JSON(http.StatusOK, user.ToResponse())
}

func (h *Handler) GetProfileById(c *gin.Context, id openapi_types.UUID) {
	repo := repository.NewUserRepository(h.DB)
	userProfile, err := repo.GetProfileWithHistory(c, id)
	if err != nil {
		slog.Info(err.Error(), "handler", "GetProfileByIdHandler")
		c.JSON(http.StatusNotFound, gin.H{"data": "User has no profile"})
		return
	}
	c.JSON(http.StatusOK, userProfile.ToResponse())
}

func (h *Handler) ListUsers(c *gin.Context) {
	repo := repository.NewUserRepository(h.DB)
	users, err := repo.List(c)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ListUsersHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.NewUsersFromModels(users))
}

func (h *Handler) UpdateUser(c *gin.Context, id openapi_types.UUID) {
	var user httpserver.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	passwordHash, err := helper.HashPassword(user.Password)
	if err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	repo := repository.NewUserRepository(h.DB)
	updateUser := &model.User{
		Id:           id,
		Username:     user.UserName,
		DisplayName:  helper.ToNullString(user.DisplayName),
		Role:         user.Role,
		Status:       user.Status,
		PasswordHash: passwordHash,
		AvatarUrl:    helper.ToNullString(user.AvatarUrl),
	}

	if err := repo.Update(c, updateUser); err != nil {
		slog.Warn(err.Error(), "handler", "UpdateUserHandler")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User updated successfully"})
}
