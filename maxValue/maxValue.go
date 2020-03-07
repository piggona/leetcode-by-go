package maxValue

func maxValue(grid [][]int) int {
	result := 0
	helper(grid, 0, 0, 0, &result)
	return result
}

func helper(grid [][]int, i, j, temp int, result *int) {
	if i == len(grid)-1 && j == len(grid[0])-1 {
		if (temp + grid[i][j]) > *result {
			*result = temp + grid[i][j]
		}
		return
	}
	if i+1 < len(grid) {
		helper(grid, i+1, j, temp+grid[i][j], result)
	}
	if j+1 < len(grid[0]) {
		helper(grid, i, j+1, temp+grid[i][j], result)
	}
	return
}
