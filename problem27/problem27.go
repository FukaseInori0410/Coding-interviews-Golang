package problem27

type BinaryTreeNode struct {
	Val   float64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Mirror(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		Mirror(root.Left)
	}
	if root.Right != nil {
		Mirror(root.Right)
	}
}
