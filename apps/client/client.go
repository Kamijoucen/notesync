package client

import (
	"context"
	"fmt"
	"time"

	"github.com/kamijoucen/notesync/apps/client/internal/ctx"
	"github.com/kamijoucen/notesync/pkg/common"
	"github.com/kamijoucen/notesync/pkg/definition"
)

func Scan(clientCtx *ctx.ClientContext, ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second):
			if err := scanFile(clientCtx); err != nil {
				// TODO: log
			}
		case <-ctx.Done():
			return
		}
	}
}

// ScanFile 扫描文件
func scanFile(clientCtx *ctx.ClientContext) error {
	allFiles, err := common.ScanLocalFiles(clientCtx.RepoPath)
	if err != nil {
		return err
	}
	for _, newFile := range allFiles {
		// TODO
		fmt.Println(newFile)
	}
	return nil
}

// syncFile 同步文件
func syncFile(clientCtx *ctx.ClientContext, f *definition.FileTreeNode) error {

	return nil
}
