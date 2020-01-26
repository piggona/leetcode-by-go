package RangeSumQuery2D

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

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
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
