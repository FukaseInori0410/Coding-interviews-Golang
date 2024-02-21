package problem21

func Reorder(arr []int) {
	rule := func(n int) bool {
		if n&0x1 == 1 {
			return true
		}
		return false
	}
	l, r := 0, len(arr)-1
	for l < r {
		for l < r && rule(arr[l]) {
			l++
		}
		for l < r && !rule(arr[r]) {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
}
