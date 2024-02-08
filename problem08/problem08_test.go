package problem08

import "testing"

func TestGetNext(t *testing.T) {
	n1 := &Node{
		Val: 1,
	}
	n2 := &Node{
		Val:    2,
		Parent: n1,
	}
	n3 := &Node{
		Val:    3,
		Parent: n1,
	}
	n1.Left = n2
	n1.Right = n3
	n4 := &Node{
		Val:    4,
		Parent: n2,
	}
	n5 := &Node{
		Val:    5,
		Parent: n2,
	}
	n2.Left = n4
	n2.Right = n5
	var testTable = []struct {
		in   *Node
		next *Node
	}{
		{nil, nil},
		{n2, n5},
		{n5, n1},
		{n3, nil},
	}
	for i, test := range testTable {
		ans := getNext(test.in)
		if ans != test.next {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
