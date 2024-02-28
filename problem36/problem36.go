package problem36

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Convert(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	var tail *BinaryTreeNode
	ConvertNode(root, &tail)
	index := root
	for index.Left != nil {
		index = index.Left
	}
	return index
}

func ConvertNode(root *BinaryTreeNode, tail **BinaryTreeNode) {
	if root == nil {
		return
	}
	index := root
	if index.Left != nil {
		ConvertNode(index.Left, tail)
	}
	index.Left = *tail
	if (*tail) != nil {
		(*tail).Right = index
	}
	*tail = index
	if index.Right != nil {
		ConvertNode(index.Right, tail)
	}
}
