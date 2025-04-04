package handler

import (
	"log/slog"
	"net/http"

	"ctf01d/internal/helper"
	"ctf01d/internal/httpserver"
	"ctf01d/internal/repository"
	"ctf01d/internal/view"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignInUser(c *gin.Context) {
	var req httpserver.SignInUserJSONBody
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn(err.Error(), "handler", "SignInUser")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetByUserName(c.Request.Context(), *req.UserName)
	if err != nil || !helper.CheckPasswordHash(*req.Password, user.PasswordHash) {
		slog.Warn(err.Error(), "handler", "SignInUser")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password or user"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	sessionId, err := repo.StoreSessionInDB(c.Request.Context(), user.Id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "SignInUser")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session"})
		return
	}

	c.SetCookie("session_id", sessionId, 96*3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"data": "User logged in"})
}

func (h *Handler) SignOutUser(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "SignOutUser")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No session found"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	if err := repo.DeleteSessionInDB(c.Request.Context(), cookie); err != nil {
		slog.Warn(err.Error(), "handler", "SignOutUser")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	// Удаление cookie
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"data": "User logout successful"})
}

func (h *Handler) ValidateSession(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No session found"})
		return
	}

	repo := repository.NewSessionRepository(h.DB)
	userId, err := repo.GetSessionFromDB(c.Request.Context(), cookie)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No user or session found"})
		return
	}

	userRepo := repository.NewUserRepository(h.DB)
	user, err := userRepo.GetById(c.Request.Context(), userId)
	if err != nil {
		slog.Warn(err.Error(), "handler", "ValidateSession")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find user by user id"})
		return
	}

	c.JSON(http.StatusOK, view.NewSessionFromModel(user))
}
