package problem35

import (
	"fmt"
	"testing"
)

func TestCloneComplexList(t *testing.T) {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3, Sibling: n5}
	n1 := &ListNode{Val: 1, Next: n2, Sibling: n3}
	n4.Sibling = n2
	c1 := CloneComplexList(n1)
	PrintList(c1)
}

func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println("List is nil!")
	}
	pivot := head
	for pivot != nil {
		fmt.Print(pivot.Val)
		pivot = pivot.Next
	}
	fmt.Print("\n")
}
