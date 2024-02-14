package problem12

import "testing"

func TestHasPath(t *testing.T) {
	var testTable = []struct {
		matrix [][]byte
		str    string
		res    bool
	}{
		{[][]byte{{'a', 'b', 't', 'g'}, {'c', 'f', 'c', 's'}, {'j', 'd', 'e', 'h'}}, "bfce", true},
		{nil, "", false},
	}
	for i, test := range testTable {
		if ans := hasPath(test.matrix, test.str); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
