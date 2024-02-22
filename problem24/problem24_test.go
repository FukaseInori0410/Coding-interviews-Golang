package problem24

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	n7 := &ListNode{5, nil}
	n6 := &ListNode{4, n7}
	n5 := &ListNode{4, n6}
	n4 := &ListNode{3, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	PrintList(ReverseList(n1))
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
