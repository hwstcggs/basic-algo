package tree

type biTreeNode struct {
	data  interface{} // 节点值
	left  *biTreeNode // 左子树
	right *biTreeNode // 右子树
}

func NewBiTree(data interface{}) *biTreeNode {
	return &biTreeNode{
		data:data,
	}
}

func (t *biTreeNode) CreateTree() *biTreeNode {

}
