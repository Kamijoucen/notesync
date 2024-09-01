package router

import (
	"github.com/labstack/echo/v4"

	"github.com/kamijoucen/notesync/internal/ctx"
	"github.com/kamijoucen/notesync/internal/services"
)

// SyncRepoMeta 创建仓库元数据
func SyncRepoMeta(serverCtx *ctx.ServerContext) func(c echo.Context) error {
	return nil
}

// FileUpload 上传单个文件
func FileUpload(serverCtx *ctx.ServerContext) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqCtx := ctx.NewRequestContext(serverCtx, c)
		fromfile, err := c.FormFile("file")
		if err != nil {
			return err
		}
		return services.UploadFile(reqCtx, fromfile)
	}
}
