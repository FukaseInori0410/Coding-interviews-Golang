package problem36

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	n7 := &BinaryTreeNode{Val: 16}
	n6 := &BinaryTreeNode{Val: 12}
	n5 := &BinaryTreeNode{Val: 8}
	n4 := &BinaryTreeNode{Val: 4}
	n3 := &BinaryTreeNode{Val: 14, Left: n6, Right: n7}
	n2 := &BinaryTreeNode{Val: 6, Left: n4, Right: n5}
	n1 := &BinaryTreeNode{Val: 10, Left: n2, Right: n3}
	PrintList(Convert(n1))

}

func PrintList(head *BinaryTreeNode) {
	if head == nil {
		fmt.Println("List is nil!")
	}
	pivot := head
	for pivot != nil {
		fmt.Print(pivot.Val)
		pivot = pivot.Right
	}
	fmt.Print("\n")
}
