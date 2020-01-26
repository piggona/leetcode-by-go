package majorityElement

func majorityElementMap(nums []int) int {
	countMap := make(map[int]int)
	count := len(nums) / 2
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]]++
		if countMap[nums[i]] >= count {
			return nums[i]
		}
	}
	return 0
}

func majorityElement(nums []int) int {
	counter := 1
	candidate := nums[0]
	for i := 1; i < len(nums); i++ {
		if counter == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			counter++
		} else {
			counter--
		}
	}
	return candidate
}
