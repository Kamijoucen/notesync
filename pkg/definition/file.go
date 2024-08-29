package definition

type FileTreeNode struct {
	Name     string
	IsDir    bool
	Hash     string
	Children []*FileTreeNode
}
