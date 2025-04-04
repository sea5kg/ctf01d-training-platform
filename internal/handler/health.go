package handler

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	version   = "dev"
	buildTime = time.Now().Format(time.RFC822Z)
)

func (h *Handler) GetVersion(c *gin.Context) {
	res := map[string]string{
		"version":    version,
		"golang":     runtime.Version(),
		"build_time": buildTime,
	}
	c.JSON(http.StatusOK, res)
}
