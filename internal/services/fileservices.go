package services

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/kamijoucen/notesync/internal/ctx"
)

func UploadFile(svc *ctx.RequestContext, fileheader *multipart.FileHeader) error {
	// joinpath
	filepath := path.Join(svc.ServerCtx.Config.FilePath, fileheader.Filename)
	// save file
	src, err := fileheader.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		_ = src.Close()
	}(src)

	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
