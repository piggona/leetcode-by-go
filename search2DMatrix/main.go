func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	targetMatrix := bs(matrix, target)
	if len(targetMatrix) == 0 {
		return false
	}
	return binarySearch(targetMatrix, target)
}

func bs(matrix [][]int, target int) []int {
	if len(matrix) == 1 {
		return matrix[0]
	}
	l := 0
	r := len(matrix) - 1

	for l <= r {
		mid := l + (r-l)/2
		cs := matrix[mid]

		if first(cs) > target {
			r = mid - 1
		} else if last(cs) >= target {
			return matrix[mid]
		} else {
			l = mid + 1
		}
	}
	return []int{}
}

func first(s []int) int {
	return s[0]
}

func last(s []int) int {
	return s[len(s)-1]
}

func binarySearch(arr []int, target int) bool {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
