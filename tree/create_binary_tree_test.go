package tree

import (
	"fmt"
	"testing"
)

func TestTraverseTree(t *testing.T) {
	treeNodes := []interface{}{"A", "B", "C",5, "D", "E"}
	tree := CreateBiTree(treeNodes,PreOrder)
	TraverseTree(tree,PreOrder)
	fmt.Println("###################")
	//fmt.Printf("tree: %+v\n",tree)
    //fmt.Printf("tree.left: %+v\n",tree.left)
	//fmt.Printf("tree.left.left: %+v\n",tree.left.left)
	tree = CreateBiTree(treeNodes,PreOrder)
	PreOrderTraverse(tree)
}