package problem29

import "fmt"

func PrintMatrixClock(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])
	if matrix == nil || rows <= 0 || columns <= 0 {
		return
	}
	start := 0
	for rows > 2*start && columns > 2*start {
		PrintMatrixCircle(matrix, rows, columns, start)
		start += 1
	}
}

func PrintMatrixCircle(matrix [][]int, rows, columns, start int) {
	endX := columns - 1 - start
	endY := rows - 1 - start
	for i := start; i <= endX; i++ {
		fmt.Printf("%d ", matrix[start][i])
	}
	if endY > start {
		for i := start + 1; i <= endY; i++ {
			fmt.Printf("%d ", matrix[i][endX])
		}
	}
	if endX > start && endY > start {
		for i := endX - 1; i >= start; i-- {
			fmt.Printf("%d ", matrix[endY][i])
		}
	}
	if endX > start && endY > start+1 {
		for i := endY - 1; i > start; i-- {
			fmt.Printf("%d ", matrix[i][start])
		}
	}
}
