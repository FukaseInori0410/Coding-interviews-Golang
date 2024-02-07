package problem05

func replaceBlank(in []byte, length int) []byte {
	if in == nil || length < len(in) {
		return in
	}
	count := 0
	for _, v := range in {
		if v == ' ' {
			count += 1
		}
	}
	oriLen := len(in)
	newLen := oriLen + 2*count
	if newLen > length {
		return in
	}
	for i := oriLen; i < newLen; i++ {
		in = append(in, 'x')
	}
	for l, r := oriLen-1, newLen-1; l < r; {
		if in[l] == ' ' {
			in[r] = '0'
			r--
			in[r] = '2'
			r--
			in[r] = '%'
			r--
			l--
		} else {
			in[r] = in[l]
			l--
			r--
		}
	}
	return in
}
