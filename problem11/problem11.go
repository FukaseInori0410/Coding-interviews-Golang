package problem11

func Min(arr []int) int {
	if arr == nil {
		return -1
	}
	l, r, mid := 0, len(arr)-1, 0

	for arr[l] >= arr[r] {
		if r-l == 1 {
			mid = r
			break
		}
		mid = (l + r) >> 1
		if arr[mid] == arr[l] && arr[mid] == arr[r] {
			mid = l
			for i := l + 1; i < r; i++ {
				if arr[i] < arr[mid] {
					mid = i
				}
			}
			break
		}
		if arr[mid] >= arr[l] {
			l = mid
		} else if arr[mid] <= arr[r] {
			r = mid
		}
	}
	return arr[mid]
}
