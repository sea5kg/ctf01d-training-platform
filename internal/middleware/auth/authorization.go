package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"ctf01d/internal/httpserver"
)

func AuthorizationMiddleware(enforcer *casbin.Enforcer) httpserver.MiddlewareFunc {
	return func(c *gin.Context) {
		// Проверка, нужен ли вообще проверять запрос на авторизацию
		_, sessionExists := c.Keys["sessionAuth.Scopes"]

		// Если оба отсутствуют, пропускаем запрос дальше
		if !sessionExists {
			c.Next()
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusForbidden, map[string]interface{}{
				"error_code": http.StatusForbidden,
				"message":    "Missing user_id",
				"details": map[string]interface{}{
					"user_id": userID,
				},
			})
			c.Abort()
			return
		}

		userIDInt, ok := userID.(int)
		if !ok {
			c.JSON(http.StatusForbidden, map[string]interface{}{
				"error_code": http.StatusForbidden,
				"message":    "Invalid user_id type",
				"details": map[string]interface{}{
					"user_id": userID,
				},
			})
			c.Abort()
			return
		}

		action := mapHTTPMethodToAction(c.Request.Method)
		if action == "" {
			c.JSON(http.StatusMethodNotAllowed, map[string]interface{}{
				"error_code": http.StatusMethodNotAllowed,
				"message":    "Unsupported HTTP method",
				"details": map[string]interface{}{
					"method": c.Request.Method,
					"action": action,
				},
			})
			c.Abort()
			return
		}

		object := getObjectFromPath(c.Request.URL.Path)
		domain := getDomain(c)

		subject := fmt.Sprintf("user_%d", userIDInt)

		allowed, err := enforcer.Enforce(subject, domain, object, action)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_code": http.StatusForbidden,
				"message":    "Authorization check error",
				"details": map[string]interface{}{
					"error": err.Error(),
				},
			})
			c.Abort()
			return
		}

		// If not allowed, return access denied
		if !allowed {
			c.JSON(http.StatusForbidden, map[string]interface{}{
				"error":   string(http.StatusForbidden),
				"message": "Access denied",
				"details": &map[string]interface{}{
					"subject": subject,
					"domain":  domain,
					"object":  object,
					"action":  action,
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func mapHTTPMethodToAction(method string) string {
	switch method {
	case http.MethodGet:
		return "view"
	case http.MethodPost:
		return "create"
	case http.MethodPut, http.MethodPatch:
		return "edit"
	case http.MethodDelete:
		return "delete"
	default:
		return ""
	}
}

func getDomain(c *gin.Context) string {
	// Check path parameters
	organizationIDStr := c.Param("organizationID")
	projectIDStr := c.Param("projectID")

	if organizationIDStr != "" {
		organizationID, err := strconv.Atoi(organizationIDStr)
		if err == nil {
			return fmt.Sprintf("organization_%d", organizationID)
		}
	}
	if projectIDStr != "" {
		projectID, err := strconv.Atoi(projectIDStr)
		if err == nil {
			return fmt.Sprintf("project_%d", projectID)
		}
	}
	// Backup body reader
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "global"
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Check request body
	var body struct {
		OrganizationID *int `json:"organization_id"`
		ProjectID      *int `json:"project_id"`
	}
	if err := json.Unmarshal(bodyBytes, &body); err == nil {
		if body.OrganizationID != nil {
			return fmt.Sprintf("organization_%d", *body.OrganizationID)
		}
		if body.ProjectID != nil {
			return fmt.Sprintf("project_%d", *body.ProjectID)
		}
	}
	return "global"
}

func getObjectFromPath(path string) string {
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) > 1 {
		return segments[1] // 'organizations', 'projects' и т.д.
	}
	return segments[0]
}
