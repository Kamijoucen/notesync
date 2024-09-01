package definition

import (
	"io"
	"unique"
)

type FileTreeNode struct {
	Flag    unique.Handle[string] // hash(path + filename) 标识唯一文件
	Name    string
	IsDir   bool
	Hash    string // 文件hash
	Path    string
	PathArr []string
}

type FileSource struct {
	Path   string
	Name   string
	Reader io.Reader
}
