package httpserver

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var tmplPath = "./html/"

func NewHtmlRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqPath := c.Request.URL.Path

		if strings.HasPrefix(reqPath, "/api/") || strings.HasSuffix(reqPath, "/api") {
			c.String(http.StatusNotFound, "API handler not found")
			return
		}

		cleanPath := filepath.Clean(filepath.Join(tmplPath, reqPath))
		fi, err := os.Stat(cleanPath)

		if (os.IsNotExist(err) || fi.IsDir()) && strings.HasPrefix(reqPath, "/assets/") {
			c.String(http.StatusNotFound, "File in assets not found")
			return
		}

		if os.IsNotExist(err) || fi.IsDir() {
			// Serve index.html для SPA
			c.File(filepath.Join(tmplPath, "index.html"))
			return
		}

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Serve static file
		http.ServeFile(c.Writer, c.Request, cleanPath)
	}
}
