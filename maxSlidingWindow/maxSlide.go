package maxSlide

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := 1
	max := []int{nums[0]}
	res := []int{}
	for right < len(nums) {
		if nums[right] <= max[0] {
			temp := nums[right]
			var i int
			for i = len(max) - 1; i >= 0; i-- {
				if temp <= max[i] {
					break
				}
			}
			max = append(max[:i+1], temp)
		} else {
			max = []int{nums[right]}
		}
		if right-left < k-1 {
			right++
		} else {
			res = append(res, max[0])
			if nums[left] == max[0] {
				max = max[1:]
			}
			left++
			right++
		}
	}
	return res
}
