package router

import (
	"github.com/kamijoucen/notesync/internal/ctx"
	"github.com/kamijoucen/notesync/internal/services"
	"github.com/labstack/echo/v4"
)

func FileUpload(serverCtx *ctx.ServerContext) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqCtx := ctx.NewRequestContext(serverCtx, c)
		return services.UploadFile(reqCtx)
	}
}
