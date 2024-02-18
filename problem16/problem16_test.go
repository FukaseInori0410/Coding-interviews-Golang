package problem16

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	res, err := Power(0.0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
