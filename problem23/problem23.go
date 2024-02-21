package problem23

type ListNode struct {
	Val  int
	Next *ListNode
}

func MeetingNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != nil || fast != nil {
		if slow == fast {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return nil
}

func EntryOfLoop(head *ListNode) *ListNode {
	mn := MeetingNode(head)
	if mn == nil {
		return nil
	}
	index := mn.Next
	count := 1
	for index != mn {
		index = index.Next
		count += 1
	}
	left, right := head, head
	for i := 0; i < count; i++ {
		right = right.Next
	}
	for left != right {
		left = left.Next
		right = right.Next
	}
	return right
}
