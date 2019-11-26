package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos1 := m - 1
	pos2 := n - 1
	pos := len(nums1) - 1
	for pos1 >= 0 && pos2 >= 0 {
		first := nums1[pos1]
		second := nums2[pos2]
		if first > second {
			nums1[pos] = first
			pos1--
		} else {
			nums1[pos] = second
			pos2--
		}
		pos--
	}
	for pos2 >= 0 {
		nums1[pos] = nums2[pos2]
		pos--
		pos2--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 7, 0, 0, 0}
	nums2 := []int{2, 3, 5, 6}
	merge(nums1, 4, nums2, 3)
	fmt.Println(nums1)
}
