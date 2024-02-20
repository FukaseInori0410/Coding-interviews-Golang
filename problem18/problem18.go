package problem18

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(head **ListNode, pivot *ListNode) {
	if *head == nil || pivot == nil {
		return
	}
	if pivot.Next != nil {
		pivot.Val = pivot.Next.Val
		pivot.Next = pivot.Next.Next
	} else if pivot == *head {
		*head = nil
		return
	} else {
		index := *head
		for index.Next != pivot {
			index = index.Next
		}
		index.Next = nil
	}
}

func DeleteDup(head **ListNode) {
	if head == nil || *head == nil || (*head).Next == nil {
		return
	}
	newHead := &ListNode{Next: *head}
	pre := newHead
	pivot := *head
	for pivot != nil {
		if pivot.Next != nil && pivot.Val == pivot.Next.Val {
			for pivot.Next != nil && pivot.Val == pivot.Next.Val {
				pivot = pivot.Next
			}
			pre.Next = pivot.Next
			pivot = pivot.Next
		} else {
			pivot = pivot.Next
			pre = pre.Next
		}
	}
	*head = newHead.Next
}
