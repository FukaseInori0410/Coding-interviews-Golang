package problem06

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintListReversingly(t *testing.T) {
	n3 := &ListNode{3, nil}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	var testTable = []*ListNode{n1}
	for _, test := range testTable {
		fmt.Println(captureOutput(func() { PrintListReversingly(test) }))
	}
}

func captureOutput(f func()) string {
	// 保存当前的标准输出
	originalOutput := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 执行函数
	f()

	// 恢复原始标准输出
	w.Close()
	os.Stdout = originalOutput

	// 读取捕获的输出
	var capturedOutput bytes.Buffer
	_, _ = io.Copy(&capturedOutput, r)

	return capturedOutput.String()
}
