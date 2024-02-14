package problem11

import "testing"

func TestMin(t *testing.T) {
	var testTable = []struct {
		arr []int
		res int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{1, 0, 1, 1, 1}, 0},
		{[]int{1, 1, 1, 0, 1}, 0},
	}
	for i, test := range testTable {
		if ans := Min(test.arr); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
