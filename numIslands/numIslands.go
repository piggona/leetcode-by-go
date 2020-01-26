package numIslands

func numIslands(grid [][]byte) int {
	res := 0
	bitmap := make([][]bool, len(grid))
	for i := 0; i < len(bitmap); i++ {
		bitmap[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				if (i == 0 || grid[i-1][j] == '0' || (grid[i-1][j] == '1' && !bitmap[i-1][j])) &&
					(i == len(grid)-1 || grid[i+1][j] == '0' || (grid[i+1][j] == '1' && !bitmap[i+1][j])) &&
					(j == 0 || grid[i][j-1] == '0' || (grid[i][j-1] == '1' && !bitmap[i][j-1])) &&
					(j == len(grid[0])-1 || grid[i][j+1] == '0' || (grid[i][j+1] == '1' && !bitmap[i][j+1])) {
					res++
				}
				bitmap[i][j] = true
			}
		}
	}
	return res
}
