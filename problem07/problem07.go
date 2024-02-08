package problem07

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func construct(preorder, inorder []int) *Node {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	rootIndex := 0
	for k, v := range inorder {
		if v == rootVal {
			rootIndex = k
		}
	}
	inL, inR := inorder[:rootIndex], inorder[rootIndex+1:]
	preL, preR := preorder[1:rootIndex+1], preorder[rootIndex+1:]
	left := construct(preL, inL)
	right := construct(preR, inR)
	return &Node{rootVal, left, right}
}
