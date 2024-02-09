package problem09

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q1 := Queue{}
	q1.AppendTail(1)
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	q1.AppendTail(2)
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	q1.DeleteHead()
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	q1.AppendTail(3)
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	q1.DeleteHead()
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	q1.DeleteHead()
	fmt.Println(q1.stack1.List)
	fmt.Println(q1.stack2.List)
	a, err := q1.DeleteHead()
	fmt.Println(a, err)
}
