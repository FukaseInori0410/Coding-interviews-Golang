package problem12

func hasPath(matrix [][]byte, str string) bool {
	if matrix == nil || str == "" {
		return false
	}
	strb := []byte(str)
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if hasPathCore(matrix, rows, cols, i, j, 0, strb) {
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix [][]byte, rows, cols, row, col, length int, strb []byte) bool {
	if row < 0 || col < 0 || row >= rows || col >= cols || matrix[row][col] != strb[length] {
		return false
	}
	if length == len(strb)-1 {
		return true
	}
	matrix[row][col] = 0x0
	res := hasPathCore(matrix, rows, cols, row-1, col, length+1,
		strb) || hasPathCore(matrix, rows, cols, row+1, col, length+1,
		strb) || hasPathCore(matrix, rows, cols, row+1, col, length+1,
		strb) || hasPathCore(matrix, rows, cols, row, col-1, length+1,
		strb) || hasPathCore(matrix, rows, cols, row, col+1, length+1, strb)
	matrix[row][col] = strb[length]
	return res
}
