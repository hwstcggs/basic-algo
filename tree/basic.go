package tree

type BiTreeNode struct {
	data  interface{} // 节点值
	left  *BiTreeNode // 左子树
	right *BiTreeNode // 右子树
}

func NewBiTreeNode(data interface{}) *BiTreeNode {
	return &BiTreeNode{
		data:data,
	}
}

func (t *BiTreeNode) CreateTree() *BiTreeNode {
	return &BiTreeNode{}
}
