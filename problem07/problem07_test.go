package problem07

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func printPreOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		printPreOrder(root.Left)
		printPreOrder(root.Right)
	}
}

func printInOrder(root *Node) {
	if root != nil {
		printInOrder(root.Left)
		fmt.Printf("%d ", root.Val)
		printInOrder(root.Right)
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

func TestConstruct(t *testing.T) {
	var testTable = []struct {
		pre []int
		in  []int
	}{
		{[]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 6, 8}},
		{nil, nil},
	}
	for _, test := range testTable {
		root := construct(test.pre, test.in)
		fmt.Println(captureOutput(func() { printPreOrder(root) }))
		fmt.Println(captureOutput(func() { printInOrder(root) }))
	}
}
