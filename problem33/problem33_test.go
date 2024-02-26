package problem33

import "testing"

func TestVerifyBST(t *testing.T) {
	var testTable = []struct {
		sequence []int
		res      bool
	}{
		{[]int{5, 7, 6, 9, 11, 10, 8}, true},
		{[]int{7, 4, 6, 5}, false},
	}
	for i, test := range testTable {
		if ans := VerifyBST(test.sequence); ans != test.res {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
