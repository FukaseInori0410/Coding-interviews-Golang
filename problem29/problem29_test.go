package problem29

import "testing"

func TestPrintMatrixClock(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	PrintMatrixClock(matrix)
}
