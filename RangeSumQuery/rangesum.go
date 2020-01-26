package RangeSumQuery

type NumArray struct {
	list []int
	dp   []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	dpArray := make([]int, len(nums))
	dpArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dpArray[i] = dpArray[i-1] + nums[i]
	}
	return NumArray{
		list: nums,
		dp:   dpArray,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.dp[j]
	}
	return this.dp[j] - this.dp[i-1]
}
