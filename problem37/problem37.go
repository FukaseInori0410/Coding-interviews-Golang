package problem37

type BinaryTreeNode struct {
	Val   rune
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Serialize(root *BinaryTreeNode) string {
	if root == nil {
		return ""
	}
	var res []rune
	var dfs func(*BinaryTreeNode)
	dfs = func(node *BinaryTreeNode) {
		if node == nil {
			res = append(res, '$')
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return string(res)
}
func Deserialize(a string, i *int) *BinaryTreeNode {
	if a[*i] == '$' {
		*i++
		return nil
	}
	root := &BinaryTreeNode{Val: rune(a[*i])}
	*i++
	root.Left = Deserialize(a, i)
	root.Right = Deserialize(a, i)
	return root
}
