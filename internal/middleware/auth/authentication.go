package auth

import (
	"net/http"

	"ctf01d/internal/httpserver"
	"ctf01d/internal/repository"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(repo repository.SessionRepository) httpserver.MiddlewareFunc {
	return func(c *gin.Context) {
		// Проверка, нужен ли вообще роуту токен для авторизации
		_, sessionRequired := c.Keys["sessionAuth.Scopes"]

		if !sessionRequired {
			c.Next()
			return
		}

		sessionCookie, err := c.Cookie("session_id")
		if err != nil || sessionCookie == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session cookie required"})
			return
		}

		userId, err := repo.GetSessionFromDB(c, sessionCookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}
