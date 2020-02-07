package findNumbers

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := []int{}
	for i := 0; i < n; i++ {
		nums[(nums[i]-1)%n] += n
	}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}
