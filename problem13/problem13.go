package problem13

func movingCount(k, rows, cols int) int {
	if rows <= 0 || cols <= 0 || k < 0 {
		return 0
	}
	res := 0
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	res = movingCountCore(k, rows, cols, 0, 0, res, visited)
	return res
}

func movingCountCore(k, rows, cols, row, col, res int, visited [][]bool) int {
	if row < 0 || col < 0 || row >= rows || col >= cols || visited[row][col] || !getSum(k, row, col) {
		return res
	}
	visited[row][col] = true
	res = 1 + movingCountCore(k, rows, cols, row+1, col, res, visited) + movingCountCore(k, rows, cols, row-1, col, res, visited) +
		movingCountCore(k, rows, cols, row, col+1, res, visited) + movingCountCore(k, rows, cols, row, col-1, res, visited)
	return res
}

func getSum(k, row, col int) bool {
	sum := 0
	for row > 0 {
		sum += row % 10
		row /= 10
	}
	for col > 0 {
		sum += col % 10
		col /= 10
	}
	if sum > k {
		return false
	}
	return true
}
