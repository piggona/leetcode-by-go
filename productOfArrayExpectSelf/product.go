package productOfArrayExpectSelf

func productExceptSelfWithDivide(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	all := 1
	for i := 0; i < len(nums); i++ {
		all = all * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = all / nums[i]
	}
	return nums
}

func productExceptSelf(nums []int) []int {
	dpFormer := make([]int, len(nums))
	dpLater := make([]int, len(nums))
	accu := 1
	dpFormer[0] = 1
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			accu = accu * nums[i]
			dpFormer[i+1] = accu
		}
	}
	accu = 1
	dpLater[len(nums)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			accu = accu * nums[i]
			dpLater[i-1] = accu
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = dpFormer[i] * dpLater[i]
	}
	return nums
}
