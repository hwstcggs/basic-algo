package tree

import "fmt"

type ElementType interface{}

type BiTreeNode struct { // 二叉链表的结点定
	data  ElementType // 节点值
	left  *BiTreeNode // 左子树
	right *BiTreeNode // 右子树
}

type TriTNode struct { // 三叉链表的结点定义
	data   ElementType //存储数据
	lChild *BiTreeNode //左孩子节点
	rChild *BiTreeNode //右孩子节点
	parent *BiTreeNode //指向双亲节点
}

func NewBiTreeNode(data ElementType) *BiTreeNode {
	return &BiTreeNode{
		data: data,
	}
}

func PreOrderTraverse(root *BiTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	if root.left == nil {
		fmt.Println("<-nil->")
	}
	PreOrderTraverse(root.left)
	if root.right == nil {
		fmt.Println("<-nil->")
	}
	PreOrderTraverse(root.right)
}

func InOrderTraverse(root *BiTreeNode) {
	if root == nil {
		return
	}
	if root.left == nil { // 打印不存在的叶子节点
		fmt.Println("<-nil->")
	}
	InOrderTraverse(root.left)
	fmt.Println(root.data)
	if root.right == nil {
		fmt.Println("<-nil->")
	}
	InOrderTraverse(root.right)
}

func PostOrderTraverse(root *BiTreeNode) {
	if root == nil {
		return
	}
	if root.left == nil {
		fmt.Println("<-nil->")
	}
	PostOrderTraverse(root.left)

	if root.right == nil {
		fmt.Println("<-nil->")
	}
	PostOrderTraverse(root.right)

	fmt.Println(root.data)
}