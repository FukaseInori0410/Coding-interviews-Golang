package problem04

func find(board [][]int, target int) bool {
	if board == nil {
		return false
	}
	row := len(board)
	col := len(board[0])
	for r, c := 0, col-1; r < row && c >= 0; {
		if board[r][c] == target {
			return true
		} else if board[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}
