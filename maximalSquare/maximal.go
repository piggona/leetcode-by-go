package maximalSquare

import "math"

type NumMatrix struct {
	data     [][]int
	dpMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{}
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for i := 1; i < len(matrix[0]); i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]
		}
	}
	return NumMatrix{
		data:     matrix,
		dpMatrix: dp,
	}

}

func (this *NumMatrix) getMatrixSum(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.data) == 0 {
		return 0
	}
	dp := this.dpMatrix
	if row1 == 0 && col1 == 0 {
		return dp[row2][col2]
	}
	if row1 == 0 {
		return dp[row2][col2] - dp[row2][col1-1]
	}
	if col1 == 0 {
		return dp[row2][col2] - dp[row1-1][col2]
	}
	return dp[row2][col2] - dp[row1-1][col2] - dp[row2][col1-1] + dp[row1-1][col1-1]
}

func maximalSquareRange(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	imatrix := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		imatrix[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			imatrix[i][j] = int(matrix[i][j]) - int('0')
		}
	}
	numMatrix := Constructor(imatrix)
	max := math.MinInt16
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if imatrix[i][j] == 1 {
				xLess := len(matrix) - i
				yLess := len(matrix[0]) - j
				maxK := getMin(xLess, yLess)
				for k := 0; k < maxK; k++ {
					sum := numMatrix.getMatrixSum(i, j, i+k, j+k)
					expect := (k + 1) * (k + 1)
					if sum == expect && sum > max {
						max = sum
					}
					if sum != expect {
						break
					}
				}
			}
		}
	}
	return max
}

func getMax(max ...int) int {
	res := math.MinInt32
	for _, value := range max {
		if value > res {
			res = value
		}
	}
	return res
}

func getMin(nums ...int) int {
	res := math.MaxInt32
	for _, value := range nums {
		if value < res {
			res = value
		}
	}
	return res
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}
	max := math.MinInt16
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				temp := getMin(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				dp[i][j] = temp + 1
				if temp+1 > max {
					max = temp
				}
			}
		}
	}
	return max
}
