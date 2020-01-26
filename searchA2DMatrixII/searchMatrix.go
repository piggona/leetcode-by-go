package searchA2DMatrixII

func searchMatrixBin(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := binarySearch(matrix, 0, target)
	col := binarySearch(matrix, 1, target)
	if row == -1 || col == -1 {
		return false
	}
	if matrix[row][0] == target || matrix[0][col] == target {
		return true
	}
	newMatrix := matrix[1 : row+1][:]
	for i := 0; i < len(newMatrix); i++ {
		temp := newMatrix[i]
		newMatrix[i] = temp[1 : col+1]
	}
	return searchMatrix(newMatrix, target)

}

func binarySearch(matrix [][]int, axis int, value int) int {
	if axis == 0 {
		n := len(matrix)
		left := 0
		right := n - 1
		for left <= right {
			middle := left + ((right - left) >> 1)
			if matrix[middle][0] > value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		if left-1 >= 0 && matrix[left-1][0] <= value {
			return left - 1
		}
		return -1
	} else {
		n := len(matrix[0])
		left := 0
		right := n - 1
		for left <= right {
			middle := left + ((right - left) >> 1)
			if matrix[0][middle] > value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		if left-1 >= 0 && matrix[0][left-1] <= value {
			return left - 1
		}
		return -1
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		temp := matrix[row][col]
		if temp < target {
			col++
		} else if temp > target {
			row--
		} else {
			return true
		}
	}
	return false
}
