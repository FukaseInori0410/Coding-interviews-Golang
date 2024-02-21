package problem23

import (
	"fmt"
	"testing"
)

func TestEntryOfLoop(t *testing.T) {
	n6 := &ListNode{6, nil}
	n5 := &ListNode{5, n6}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	n6.Next = n3
	fmt.Println(EntryOfLoop(n1))
}
