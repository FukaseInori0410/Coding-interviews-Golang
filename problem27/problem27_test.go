package problem27

import (
	"fmt"
	"testing"
)

func TestMirror(t *testing.T) {
	n3 := &BinaryTreeNode{Val: 10, Left: nil, Right: nil}
	n2 := &BinaryTreeNode{Val: 6, Left: nil, Right: nil}
	n1 := &BinaryTreeNode{Val: 8, Left: n2, Right: n3}
	Mirror(n1)
	fmt.Println(n1.Left, n1.Right)
}
