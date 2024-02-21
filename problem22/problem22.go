package problem22

type ListNode struct {
	Val  int
	Next *ListNode
}

func LastK(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return nil
	}
	fast := head
	for i := 1; i < k; i++ {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			return nil
		}
	}
	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
