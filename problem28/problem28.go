package problem28

type BinaryTreeNode struct {
	Val   float64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func IsSym(root *BinaryTreeNode) bool {
	return Sym(root, root)
}

func Sym(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return Sym(root1.Left, root2.Right) && Sym(root1.Right, root2.Left)
}
