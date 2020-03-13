package tree

import (
	"fmt"
	"testing"
)

func TestTraverseTree(t *testing.T) {
	nodes01:= []interface{}{"A", "B", "C", 5, "<-nil->", "D","<-nil->", "E"}
	//nodes02 := []interface{}{"A", "B", "C"}
	fmt.Println("#########******  start  ******##########")
	tree := CreateBiTree(nodes01)
	PreOrderTraverse(tree)
	fmt.Println("#########******  end  ******##########")
}
