package problem03

func duplicate(a []int) []int {
	var res []int
	if len(a) == 0 || len(a) == 1 {
		return nil
	}
	for i := 0; i < len(a); i++ {
		if a[i] == i {
			continue
		}
		if a[a[i]] == a[i] {
			res = append(res, a[i])
			continue
		} else {
			a[a[i]], a[i] = a[i], a[a[i]]
		}
	}
	return res
}

func getDuplication(a []int) int {
	if len(a) == 0 || len(a) == 1 {
		return -1
	}
	left, right := 1, len(a)
	for left <= right {
		mid := (right + left) >> 1
		count := func(l, r int) int {
			cnt := 0
			for _, v := range a {
				if v >= l && v <= r {
					cnt++
				}
			}
			return cnt
		}(left, mid)
		if left == right {
			if count > 1 {
				return left
			} else {
				break
			}
		}
		if count > mid-left+1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 0
}
