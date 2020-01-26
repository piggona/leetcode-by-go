package moveZeros

import "sort"

func moveZeros(nums []int) {
	left := 0
	right := len(nums) - 1
	for nums[right] == 0 {
		right--
	}
	for left < right {
		if nums[left] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			left++
		}
	}
	sort.Ints(nums[:left])
}
