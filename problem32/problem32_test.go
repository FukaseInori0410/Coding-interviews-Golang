package problem32

import "testing"

func TestPrintFromTop(t *testing.T) {
	n7 := &BinaryTreeNode{Val: 7}
	n6 := &BinaryTreeNode{Val: 4}
	n5 := &BinaryTreeNode{Val: 2, Left: n6, Right: n7}
	n4 := &BinaryTreeNode{Val: 9}
	n3 := &BinaryTreeNode{Val: 8, Left: n4, Right: n5}
	n2 := &BinaryTreeNode{Val: 7}
	n1 := &BinaryTreeNode{Val: 8, Left: n3, Right: n2}
	PrintFromTop(n1)
	PrintInRow(n1)
}
