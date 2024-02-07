package problem04

import "testing"

func TestSearch(t *testing.T) {
	var testTable = []struct {
		board  [][]int
		target int
		res    bool
	}{
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 4, true},
		{nil, 1, false},
	}
	for i, test := range testTable {
		if ans := find(test.board, test.target); ans != test.res {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
