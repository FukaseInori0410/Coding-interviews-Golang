package problem20

import "testing"

func TestIsNum(t *testing.T) {
	var testTable = []struct {
		str string
		res bool
	}{
		{"123", true},
		{"+123.45e6", true},
		{"123.", true},
		{".456", true},
		{"1a3.14", false},
	}
	for i, test := range testTable {
		if ans := isNum(test.str); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
