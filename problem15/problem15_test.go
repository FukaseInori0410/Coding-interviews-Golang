package problem15

import "testing"

func TestNumberOf1(t *testing.T) {
	var testTable = []struct {
		num   int
		count int
	}{
		{9, 2},
		{3, 2},
		{2, 1},
		{0, 0},
	}
	for i, test := range testTable {
		if ans := NumberOf1(test.num); ans != test.count {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
