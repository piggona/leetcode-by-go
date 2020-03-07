package shortestsubarray

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1
	for {
		for i := right; i > left; i-- {
			if nums[i] < nums[left] {
				goto LOOP
			}
		}
		left++
	}
LOOP:
	for {
		for i := left; i < right; i++ {
			if nums[i] > nums[right] {
				goto END
			}
		}
		right--
	}
END:
	return right - left + 1

}
