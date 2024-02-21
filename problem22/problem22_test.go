package problem22

import (
	"fmt"
	"testing"
)

func TestLastK(t *testing.T) {
	n7 := &ListNode{5, nil}
	n6 := &ListNode{4, n7}
	n5 := &ListNode{4, n6}
	n4 := &ListNode{3, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	fmt.Println(LastK(n1, 3))
	fmt.Println(LastK(n1, 8))
	fmt.Println(LastK(nil, 0))
}
