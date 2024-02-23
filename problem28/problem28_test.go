package problem28

import (
	"fmt"
	"testing"
)

func TestIsSym(t *testing.T) {
	n4 := &BinaryTreeNode{Val: 2}
	n3 := &BinaryTreeNode{Val: 2, Left: n4}
	n2 := &BinaryTreeNode{Val: 2}
	n1 := &BinaryTreeNode{Val: 2, Left: n2, Right: n3}
	fmt.Println(IsSym(n1))
}
