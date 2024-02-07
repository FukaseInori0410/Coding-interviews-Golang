package problem05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	var testTable = []struct {
		in     []byte
		length int
		res    []byte
	}{
		{nil, 0, nil},
		{append(make([]byte, 0, 14), []byte("Hello World!")...), 14, []byte("Hello%20World!")},
		{append(make([]byte, 0, 13), []byte("Hello World!")...), 13, []byte("Hello World!")},
		{append(make([]byte, 0, 11), []byte("HelloWorld!")...), 11, []byte("HelloWorld!")},
	}
	for i, test := range testTable {
		ans := replaceBlank(test.in, test.length)
		if !assert.Equal(t, test.res, ans) {
			t.Errorf("第%d个测试用例出错", i+1)
		}
	}
}
