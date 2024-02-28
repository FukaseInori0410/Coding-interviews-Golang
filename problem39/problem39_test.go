package problem39

import (
	"testing"
)

func TestMoreThanHalfNumber(t *testing.T) {
	var testTable = []struct {
		in  []int
		res int
	}{
		{[]int{2, 1, 2, 3, 2, 4, 2, 5, 2, 6, 2}, 2},
		{[]int{}, -1},
	}
	for i, test := range testTable {
		if ans := MoreThanHalfNumber(test.in); ans != test.res {
			t.Errorf("第%d个用例出错", i+1)
		}
	}
}
