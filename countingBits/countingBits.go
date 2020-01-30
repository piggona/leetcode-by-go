package countingBits

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	result := make([]int, num+1)
	result[0] = 0
	if num >= 1 {
		result[1] = 1
	}
	if num == 1 {
		return result
	}
	count := 1
	i := 2
	for {
		count = count << 2
		for j := 0; j < count; j++ {
			if i == num {
				result[i] = result[i-count] + 1
				return result
			} else {
				result[i] = result[i-count] + 1
				i++
			}
		}
	}
}
