package router

import (
	"github.com/kamijoucen/notesync/apps/server/internal/ctx"
	"github.com/kamijoucen/notesync/pkg/config"
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, cfg *config.ServerConfig) {
	serverCtx := ctx.NewServerContext(cfg)
	e.POST("/upload", FileUpload(serverCtx))
}
