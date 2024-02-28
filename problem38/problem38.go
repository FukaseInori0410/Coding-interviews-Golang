package problem38

import "fmt"

func Permutation(s string) {
	if s == "" {
		fmt.Println("nil string")
		return
	}
	index := 0
	Permute(s, index)
}

func Permute(in string, index int) {
	if index == len(in) {
		fmt.Println(in)
		return
	}
	inByte := []byte(in)
	for i := index; i < len(in); i++ {
		inByte[index], inByte[i] = inByte[i], inByte[index]
		Permute(string(inByte), index+1)
	}
}
