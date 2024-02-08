package problem08

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func getNext(node *Node) *Node {
	next := node
	if node == nil {
		return next
	}
	if node.Right != nil {
		next = node.Right
		for next.Left != nil {
			next = next.Left
		}
	} else if node == node.Parent.Left {
		next = node.Parent
	} else {
		next = node.Parent
		for next.Parent != nil && next != next.Parent.Left {
			next = next.Parent
		}
		next = next.Parent
	}
	return next
}
