package problem03

import "testing"

func TestDuplicate(t *testing.T) {
	var testTable = []struct {
		a   []int
		res []int
	}{
		{nil, nil},
		{[]int{0, 1, 2, 3}, nil},
		{[]int{3, 2, 1, 0}, nil},
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}},
	}
	for i, test := range testTable {
		ans := duplicate(test.a)
		if len(ans) != len(test.res) {
			t.Error("length unequal!")
		}
		for k, v := range ans {
			if v != test.res[k] {
				t.Errorf("第%d个测试用例的第%d个结果错误", i+1, k+1)
			}
		}
	}
}

func TestDuplicateLocal(t *testing.T) {
	var testTable = []struct {
		a   []int
		res int
	}{
		{nil, -1},
		{[]int{1, 1}, 1},
		{[]int{2, 3, 5, 4, 4, 2, 6, 7}, 4},
	}
	for i, test := range testTable {
		if ans := getDuplication(test.a); ans != test.res {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
