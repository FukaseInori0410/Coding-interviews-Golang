package problem30

import (
	"fmt"
	"testing"
)

func TestStackWithMin_Min(t *testing.T) {
	var stack StackWithMin
	stack.Push(3)
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Push(4)
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Push(2)
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Push(1)
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Pop()
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Pop()
	fmt.Println(stack.StackData, stack.StackMin)
	stack.Push(0)
	fmt.Println(stack.StackData, stack.StackMin)
}
