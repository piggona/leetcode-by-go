package findKthLargest

import "container/heap"

type hNum []int

func (h hNum) Len() int {
	return len(h)
}

func (h hNum) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hNum) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *hNum) Push(x interface{}) {
	if num, ok := x.(int); ok {
		*h = append(*h, num)
	}
}

func (h *hNum) Pop() interface{} {
	n := len(*h)
	temp := (*h)[n-1]
	*h = (*h)[:n-1]
	return temp
}

func findKthLargestHeap(nums []int, k int) int {
	h := hNum(nums)
	heap.Init(&h)
	temp := 0
	for i := 0; i < k; i++ {
		temp = (h.Pop()).(int)
	}
	return temp
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	left := 1
	right := len(nums) - 1
	for {
		if left < right && nums[right] < nums[0] {
			right--
		}
		if left < right && nums[left] > nums[0] {
			left++
		}
		if left >= right {
			break
		}
		nums[right], nums[left] = nums[left], nums[right]
	}
	nums[left], nums[0] = nums[0], nums[left]
	if left+1 == k {
		return nums[left]
	}
	if left+1 < k {
		return findKthLargest(nums[left+1:], k-left-1)
	}
	return findKthLargest(nums[:left], k)
}
