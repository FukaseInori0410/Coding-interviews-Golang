package problem35

type ListNode struct {
	Val     int
	Next    *ListNode
	Sibling *ListNode
}

func CloneComplexList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	CloneNodes(head)
	ConnectSibling(head)
	return BreakLists(head)
}
func CloneNodes(head *ListNode) {
	if head == nil {
		return
	}
	index := head
	for index != nil {
		temp := &ListNode{Val: index.Val, Next: index.Next}
		index.Next = temp
		index = temp.Next
	}
}
func ConnectSibling(head *ListNode) {
	if head == nil {
		return
	}
	index1 := head
	for index1 != nil {
		index2 := index1.Next
		if index1.Sibling != nil {
			index2.Sibling = index1.Sibling.Next
		}
		index1 = index2.Next
	}
}
func BreakLists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	headClone := head.Next
	index := head
	indexClone := head.Next
	for {
		index.Next = indexClone.Next
		if index.Next == nil {
			break
		}
		index = index.Next
		indexClone.Next = index.Next
		indexClone = indexClone.Next
	}
	return headClone
}
