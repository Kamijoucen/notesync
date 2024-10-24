package logic

import (
	"io"
	"os"
	"path"

	"github.com/kamijoucen/notesync/apps/server/internal/ctx"
	"github.com/kamijoucen/notesync/apps/server/internal/param"
	"github.com/kamijoucen/notesync/pkg/definition"
)

// SyncRepoMeta 创建仓库
func SyncRepoMeta(svc *ctx.RequestContext, req *param.CreateRepositoryRequest) (*param.CreateRepositoryResponse, error) {

	err := svc.ServerCtx.Ent.Repository.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		Exec(svc.Ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// StoreFile 向仓库存储文件
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
