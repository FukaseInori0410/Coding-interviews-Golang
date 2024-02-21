package problem21

import (
	"fmt"
	"testing"
)

func TestReorder(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	Reorder(arr)
	fmt.Println(arr)
}
