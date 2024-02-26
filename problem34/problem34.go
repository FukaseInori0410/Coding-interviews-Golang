package problem34

import "fmt"

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func PrintPath(root *BinaryTreeNode, target int) {
	if root == nil {
		return
	}
	var path []int
	sum := 0
	FindPath(root, target, sum, path)
}

func FindPath(root *BinaryTreeNode, target, sum int, path []int) {
	sum += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == target {
			fmt.Println(path)
		} else {
			return
		}
	}
	if root.Left != nil {
		FindPath(root.Left, target, sum, path)
	}
	if root.Right != nil {
		FindPath(root.Right, target, sum, path)
	}
}
