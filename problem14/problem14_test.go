package problem14

import "testing"

func TestMaxCuttingDp(t *testing.T) {
	var testTable = []struct {
		length int
		res    int
	}{
		{8, 18},
		{4, 4},
	}
	for i, test := range testTable {
		if ans := maxCuttingDp(test.length); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}

func TestMaxCuttingGreedy(t *testing.T) {
	var testTable = []struct {
		length int
		res    int
	}{
		{8, 18},
		{4, 4},
	}
	for i, test := range testTable {
		if ans := maxCuttingGreedy(test.length); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
