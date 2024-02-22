package problem25

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var newHead *ListNode
	if head1.Val <= head2.Val {
		newHead = head1
		newHead.Next = Merge(head1.Next, head2)
	} else {
		newHead = head2
		newHead.Next = Merge(head1, head2.Next)
	}
	return newHead
}
