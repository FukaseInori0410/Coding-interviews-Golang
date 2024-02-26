package problem33

type BinaryTreeNode struct {
	Val   float64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func VerifyBST(sequence []int) bool {
	if sequence == nil {
		return false
	}
	length := len(sequence)
	root := sequence[length-1]
	i := 0
	for ; i < length-1; i++ {
		if sequence[i] > root {
			break
		}
	}
	j := i
	for ; j < length-1; j++ {
		if sequence[j] < root {
			return false
		}
	}
	left, right := true, true
	if i > 0 {
		left = VerifyBST(sequence[0:i])
	}
	if i < j {
		right = VerifyBST(sequence[i:j])
	}
	return left && right
}
