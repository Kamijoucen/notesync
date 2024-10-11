package logic

import (
	"io"
	"os"
	"path"

	"github.com/kamijoucen/notesync/internal/ctx"
	"github.com/kamijoucen/notesync/pkg/definition"
)

func StoreFile(svc *ctx.RequestContext, src *definition.FileSource) error {
	// joinpath
	filepath := path.Join(svc.ServerCtx.Config.FilePath, src.Name)
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(dst *os.File) {
		_ = dst.Close()
	}(dst)
	if _, err = io.Copy(dst, src.Reader); err != nil {
		return err
	}
	return nil
}
