package problem20

func isNum(str string) bool {
	if str == "" {
		return false
	}
	byteStr := []byte(str)
	length := len(byteStr)
	res := false
	i := 0
	res, i = scanInt(byteStr, i, length)
	if i < length && byteStr[i] == '.' {
		i++
		scanRes, j := scanUsInt(byteStr, i, length)
		res = res || scanRes
		i = j
	}
	if i < length && (byteStr[i] == 'e' || byteStr[i] == 'E') {
		i++
		scanRes, j := scanInt(byteStr, i, length)
		res = res && scanRes
		i = j
	}
	return res && (i == length)
}

func scanInt(byteStr []byte, in, length int) (res bool, out int) {
	if in < length && (byteStr[in] == '+' || byteStr[in] == '-') {
		in++
	}
	return scanUsInt(byteStr, in, length)
}

func scanUsInt(byteStr []byte, in, length int) (res bool, out int) {
	index := in
	for index < length && byteStr[index] >= '0' && byteStr[index] <= '9' {
		index++
	}
	return index > in, index
}
