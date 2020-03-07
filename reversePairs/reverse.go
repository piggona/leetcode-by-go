package reversePairs

func reversePairs(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
		count int
	)
	depart(nums, left, right, &count)
	return count
}

func depart(nums []int, l, r int, count *int) {
	if l == r {
		return
	}
	mid := (l + (r-l)>>1) + 1
	depart(nums, l, mid-1, count)
	depart(nums, mid, r, count)
	merge(nums, l, mid, r, count)
}

func merge(nums []int, l, m, r int, count *int) {
	var (
		left  = make([]int, m-l)
		right = make([]int, r-m+1)
		lcur  int
		rcur  int
		pos   = l
	)
	copy(left, nums[l:m])
	copy(right, nums[m:r+1])
	for lcur < len(left) && rcur < len(right) {
		if left[lcur] <= right[rcur] {
			nums[pos] = left[lcur]
			*count += rcur
			lcur++
		} else {
			nums[pos] = right[rcur]
			rcur++
		}
		pos++
	}
	for lcur < len(left) {
		*count += rcur
		nums[pos] = left[lcur]
		lcur++
		pos++
	}
	for rcur < len(right) {
		nums[pos] = right[rcur]
		rcur++
		pos++
	}
	return
}
