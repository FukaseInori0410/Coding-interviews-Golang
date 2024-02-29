package problem40

func Partition(arr []int, start, end int) int {
	pivot := arr[start]
	l, r := start+1, end
	for {
		for l <= r && pivot > arr[l] {
			l++
		}
		for l <= r && pivot < arr[r] {
			r--
		}
		if l > r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	arr[start], arr[r] = arr[r], arr[start]
	return r
}

func On(in []int, k int) []int {
	if len(in) == 0 || k > len(in) {
		return nil
	}
	var res []int
	start := 0
	end := len(in) - 1
	index := Partition(in, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = Partition(in, start, end)
		} else {
			start = index + 1
			index = Partition(in, start, end)
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, in[i])
	}
	return res
}
