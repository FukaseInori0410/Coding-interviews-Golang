package problem37

import (
	"fmt"
	"testing"
)

func TestSerialize(t *testing.T) {
	n6 := &BinaryTreeNode{Val: '6'}
	n5 := &BinaryTreeNode{Val: '5'}
	n4 := &BinaryTreeNode{Val: '4'}
	n3 := &BinaryTreeNode{Val: '3', Left: n5, Right: n6}
	n2 := &BinaryTreeNode{Val: '2', Left: n4}
	n1 := &BinaryTreeNode{Val: '1', Left: n2, Right: n3}
	s := Serialize(n1)
	fmt.Println(s)
	i := 0
	m1 := Deserialize(s, &i)
	fmt.Println(string(m1.Right.Right.Val))
}
