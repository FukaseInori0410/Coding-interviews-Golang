package problem17

import (
	"errors"
	"fmt"
)

func Print1toMax(n int) {
	if n <= 0 {
		fmt.Println(errors.New("n<=0"))
		return
	}
	res := make([]rune, n)
	for i := 0; i < 10; i++ {
		res[0] = '0' + rune(i)
		PrintRe(res, n, 0)
	}
}

func PrintRe(num []rune, length, index int) {
	if index == length-1 {
		PrintNum(num)
		return
	}
	for i := 0; i < 10; i++ {
		num[index+1] = '0' + rune(i)
		PrintRe(num, length, index+1)
	}
}

func PrintNum(num []rune) {
	var start bool
	for i := 0; i < len(num); i++ {
		if !start && num[i] != '0' {
			start = true
		}
		if start {
			fmt.Printf("%c", num[i])
		}
	}
	fmt.Print(" ")
}
