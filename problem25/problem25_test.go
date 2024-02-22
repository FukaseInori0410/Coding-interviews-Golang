package problem25

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	n7 := &ListNode{5, nil}
	n6 := &ListNode{4, nil}
	n5 := &ListNode{4, n7}
	n4 := &ListNode{3, n6}
	n3 := &ListNode{3, n5}
	n2 := &ListNode{2, n4}
	n1 := &ListNode{1, n3}
	PrintList(n1)
	PrintList(n2)
	PrintList(Merge(n1, n2))
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
