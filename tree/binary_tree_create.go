package tree

const (
	PreOrder  = iota //先序遍历
	InOrder          //中序遍历
	PostOrder        //后序遍历
)

// 创建一棵二叉树
var count = -1
func CreateBiTree(nodes []interface{}) *BiTreeNode {
	count = count + 1
	if count >= len(nodes) {
		return nil
	}
	var node *BiTreeNode

	if nodes[count] == "<-nil->" {
		return nil
	} else {
		node = NewBiTreeNode(nodes[count])
		node.left = CreateBiTree(nodes)
		node.right = CreateBiTree(nodes)
	}
	return node
}


