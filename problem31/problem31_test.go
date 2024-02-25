package problem31

import "testing"

func TestIsPopOrder(t *testing.T) {
	var testTable = []struct {
		push   []int
		pop    []int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	}
	for i, test := range testTable {
		if ans := IsPopOrder(test.push, test.pop); ans != test.result {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
