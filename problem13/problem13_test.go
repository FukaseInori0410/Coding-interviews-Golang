package problem13

import "testing"

func TestMovingCount(t *testing.T) {
	var testTable = []struct {
		rows int
		cols int
		k    int
		res  int
	}{
		{1, 1, 0, 1},
		{1, 1, -1, 0},
		{3, 3, 2, 6},
	}
	for i, test := range testTable {
		if ans := movingCount(test.k, test.rows, test.cols); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
