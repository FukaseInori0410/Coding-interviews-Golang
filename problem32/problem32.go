package problem32

import "fmt"

type BinaryTreeNode struct {
	Val   float64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func PrintFromTop(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		fmt.Print(node.Val)
		queue = queue[1:]
	}
	fmt.Print("\n")
}

func PrintInRow(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			fmt.Print(queue[i].Val)
		}
		fmt.Print("\n")
		queue = queue[l:]
	}
}
