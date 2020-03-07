package getleast

func getLeastNumbers(arr []int, k int) []int {
	left := 1
	right := len(arr) - 1
	can := arr[0]
	for {
		for right >= 1 && arr[right] > can {
			right--
		}
		for left < len(arr) && arr[left] <= can {
			left++
		}
		if right <= left {
			break
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[0], arr[right] = arr[right], arr[0]
	if right+1 == k {
		return arr[:right+1]
	}
	if right+1 < k {
		temp := getLeastNumbers(arr[right+1:], k-right-1)
		return append(arr[:right+1], temp...)
	}
	return getLeastNumbers(arr[:right+1], k)
}

func leastNum(nums []int) []int {
	left := 0
	right := len(nums) - 1
	cur := nums[0]
	for {
		for right >= 0 && nums[right] > cur {
			right--
		}
		for left < len(nums) && nums[left] <= cur {
			left++
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[right], nums[0] = nums[0], nums[right]
	leastNum(nums[:right+1])
	leastNum(nums[right+1:])
	return nums
}
