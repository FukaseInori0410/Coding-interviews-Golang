package problem34

import "testing"

func TestPrintPath(t *testing.T) {
	n5 := &BinaryTreeNode{Val: 7}
	n4 := &BinaryTreeNode{Val: 4}
	n3 := &BinaryTreeNode{Val: 12}
	n2 := &BinaryTreeNode{Val: 5, Left: n4, Right: n5}
	n1 := &BinaryTreeNode{Val: 10, Left: n2, Right: n3}
	PrintPath(n1, 22)
}
