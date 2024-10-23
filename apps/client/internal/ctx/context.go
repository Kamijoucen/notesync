package ctx

import (
	"unique"

	"github.com/kamijoucen/notesync/pkg/definition"
)

type ClientContext struct {
	RepoPath  string
	FlagCache map[unique.Handle[string]]*definition.FileTreeNode
}

func NewClientContext(repoPath string) *ClientContext {
	return &ClientContext{
		RepoPath: repoPath,
	}
}
