package problem26

type BinaryTreeNode struct {
	Val   float64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func HaveRoot2(root1, root2 *BinaryTreeNode) bool {
	result := false
	if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			result = HaveTree2(root1, root2)
		}
		if !result {
			result = HaveRoot2(root1.Left, root2)
		}
		if !result {
			result = HaveRoot2(root1.Right, root2)
		}
	}
	return result
}

func HaveTree2(root1, root2 *BinaryTreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return HaveTree2(root1.Left, root2.Left) && HaveTree2(root1.Right, root2.Right)

}
