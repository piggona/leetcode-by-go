package straight

func isStraight(nums []int) bool {
	bitmap := make([]bool, 14)
	start := 14
	end := 0
	for _, val := range nums {
		if bitmap[val] {
			return false
		}
		if !bitmap[val] && val != 0 {
			bitmap[val] = true
			if val < start {
				start = val
			}
			if val > end {
				end = val
			}
		}
	}
	if end-start > 4 {
		return false
	}
	return true
}
