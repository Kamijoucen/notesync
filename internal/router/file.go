package router

import (
	"errors"
	"mime/multipart"

	"github.com/labstack/echo/v4"

	"github.com/kamijoucen/notesync/internal/ctx"
	"github.com/kamijoucen/notesync/internal/services"
	"github.com/kamijoucen/notesync/pkg/definition"
)

// SyncRepoMeta 创建仓库元数据
func SyncRepoMeta(serverCtx *ctx.ServerContext) func(c echo.Context) error {
	return nil
}

// FileUpload 上传单个文件
func FileUpload(serverCtx *ctx.ServerContext) func(c echo.Context) error {
	return func(c echo.Context) error {
		fromfile, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := fromfile.Open()
		if err != nil {
			return err
		}
		defer func(src multipart.File) {
			_ = src.Close()
		}(src)
		path := c.FormValue("path")

		if path == "" {
			return errors.New("path is empty")
		}
		reqCtx := ctx.NewRequestContext(serverCtx, c)
		return services.StoreFile(reqCtx, &definition.FileSource{
			Path:   path,
			Name:   fromfile.Filename,
			Reader: src,
		})
	}
}
