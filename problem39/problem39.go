package problem39

func MoreThanHalfNumber(in []int) int {
	if len(in) == 0 {
		return -1
	}
	length := len(in)
	times := 1
	res := in[0]
	for i := 1; i < length; i++ {
		if in[i] == res {
			times++
		} else {
			times--
		}
		if times == 0 {
			res = in[i]
			times = 1
		}
	}
	return res
}
