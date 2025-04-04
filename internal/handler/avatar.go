package handler

import (
	"log/slog"

	"ctf01d/internal/httpserver"
	"ctf01d/pkg/avatar"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UniqueAvatar(c *gin.Context, username string, params httpserver.UniqueAvatarParams) {
	xMax := 100
	yMax := 100
	if params.Max != nil {
		xMax = *params.Max
		yMax = *params.Max
	}
	blockSize := 20
	if params.BlockSize != nil {
		blockSize = *params.BlockSize
	}
	steps := 8
	if params.Steps != nil {
		steps = *params.Steps
	}

	imageBytes := avatar.GenerateAvatar(username, xMax, yMax, blockSize, steps)

	c.Writer.Header().Set("Content-Type", "image/png")
	if _, err := c.Writer.Write(imageBytes); err != nil {
		slog.Warn(err.Error(), "handler", "UniqueAvatar")
	}
}
