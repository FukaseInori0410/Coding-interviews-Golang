package problem14

import "math"

func maxCuttingDp(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= length; i++ {
		for j := 1; j <= i/2; j++ {
			temp := dp[j] * dp[i-j]
			if temp > dp[i] {
				dp[i] = temp
			}
		}
	}
	return dp[length]
}

func maxCuttingGreedy(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	timesOf3 := length / 3
	timesOf2 := 0
	if length-3*timesOf3 == 1 {
		timesOf3 -= 1
		timesOf2 = 2
	} else if length-3*timesOf3 == 2 {
		timesOf2 = 1
	}
	return int(math.Pow(3, float64(timesOf3))) * int(math.Pow(2, float64(timesOf2)))
}
