package tree

import "fmt"

const (
	PreOrder  = iota //先序遍历
	InOrder         //中序遍历
	PostOrder       //后序遍历
)

// 创建一棵二叉树
var count = -1
func CreateBiTree(nodes []interface{},TraverseMethod int) *BiTreeNode {
	count = count + 1
	if count >= len(nodes) {
		return nil
	}
	var node *BiTreeNode

	if nodes[count] == '#' {
		return nil
	} else {
		node = new(BiTreeNode)
		switch TraverseMethod {
		case PreOrder:
			node.data = nodes[count]
			node.left = CreateBiTree(nodes,TraverseMethod)
			node.right = CreateBiTree(nodes,TraverseMethod)
		case InOrder:
			node.left = CreateBiTree(nodes,TraverseMethod)
			node.data = nodes[count]
			node.right = CreateBiTree(nodes,TraverseMethod)
		case PostOrder:
			node.left = CreateBiTree(nodes,TraverseMethod)
			node.right = CreateBiTree(nodes,TraverseMethod)
			node.data = nodes[count]
		}
	}
	return node
}

//遍历二叉树,三种遍历法, 先序,中序,后序
func TraverseTree(root *BiTreeNode, TraverseMethod int) {
	if root == nil {
		return
	}
	switch TraverseMethod {
	case PreOrder: // 先序遍历
		fmt.Println(root.data) // 第一步取根节点

		if root.left == nil {
			fmt.Println("#")
		}
		TraverseTree(root.left, PreOrder) // 第二步取左子树节点

		if root.right == nil {
			fmt.Println("#")
		}
		TraverseTree(root.right, PreOrder) // 第三步取右子树节点

	case InOrder: // 中序遍历
		if root.left == nil {
			fmt.Println("#")
		}
		TraverseTree(root.left, InOrder) // 第一步取左子树节点

		fmt.Println(root.data) // 第二步取根节点

		if root.right == nil {
			fmt.Println("#")
		}
		TraverseTree(root.right, InOrder) // 第三步取右子树节点

	case PostOrder: // 后序遍历
		if root.left == nil {
			fmt.Println("#")
		}
		TraverseTree(root.left, PostOrder) // 第一步取左子树节点

		if root.right == nil {
			fmt.Println("#")
		}
		TraverseTree(root.right, PostOrder) // 第二步取右子树节点

		fmt.Println(root.data) // 第三步取根节点
	}
}

func PreOrderTraverse(root *BiTreeNode)  {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	if root.left == nil {
		fmt.Println("#")
	}
	PreOrderTraverse(root.left)
	if root.right == nil {
		fmt.Println("#")
	}
	PreOrderTraverse(root.right)
}

func InOrderTraverse(root *BiTreeNode)  {
	if root == nil {
		return
	}
	if root.left == nil {
		fmt.Println("#")
	}
	InOrderTraverse(root.left)
	fmt.Println(root.data)
	if root.right == nil {
		fmt.Println("#")
	}
	InOrderTraverse(root.right)
}

func PostOrderTraverse(root *BiTreeNode)  {
	if root == nil {
		return
	}
	if root.left == nil {
		fmt.Println("#")
	}
	PostOrderTraverse(root.left)

	if root.right == nil {
		fmt.Println("#")
	}
	PostOrderTraverse(root.right)

	fmt.Println(root.data)
}