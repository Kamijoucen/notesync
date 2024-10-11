package logic

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/kamijoucen/notesync/internal/ctx"
	"github.com/kamijoucen/notesync/pkg/definition"
)

func PushFile(clientCtx *ctx.ClientContext, fNode *definition.FileTreeNode) error {
	file, err := os.Open(fNode.Name)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// buffer
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// 创建表单提交文件
	part, err := writer.CreateFormFile("file", fNode.Name)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	defer func(writer *multipart.Writer) {
		_ = writer.Close()
	}(writer)
	return nil
}
