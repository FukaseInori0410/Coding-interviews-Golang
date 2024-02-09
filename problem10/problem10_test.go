package problem10

import "testing"

func TestFibonacci(t *testing.T) {
	var testTable = []struct {
		n   int
		res int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{13, 233},
	}
	for i, test := range testTable {
		if ans := Fibonacci(test.n); ans != test.res {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
