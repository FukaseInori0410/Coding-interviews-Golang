package problem19

import "testing"

func TestMatch(t *testing.T) {
	var testTable = []struct {
		str     string
		pattern string
		res     bool
	}{
		{"aaa", "a.a", true},
		{"aaa", "ab*ac*.", true},
		{"aaa", "a.b", false},
	}
	for i, test := range testTable {
		if ans := Match(test.str, test.pattern); ans != test.res {
			t.Errorf("第%d个测试用例出错，输出%v", i+1, ans)
		}
	}
}
