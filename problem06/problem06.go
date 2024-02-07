package problem06

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListReversingly(root *ListNode) {
	if root == nil {
		return
	}
	if root.Next != nil {
		PrintListReversingly(root.Next)
	}
	fmt.Print(root.Val)
}
