package problem24

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	index := head
	var prev *ListNode
	var next *ListNode
	var ReverseHead *ListNode
	for index != nil {
		next = index.Next
		if next == nil {
			ReverseHead = index
		}
		index.Next = prev
		prev = index
		index = next
	}
	return ReverseHead
}
