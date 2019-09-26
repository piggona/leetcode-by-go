package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	pos := (len(nums1) + len(nums2)) / 2
	first := 0
	second := 0
	result := 0
	pre := 0
	if pos == 0 {
		if len(nums1) != 0 {
			return float64(nums1[0])
		}
		return float64(nums2[0])
	}
	for i := 0; i < pos+1; i++ {
		if first >= len(nums1) {
			cur := second + pos - i
			pre = result
			result = nums2[cur]
			if cur != 0 && pre < nums2[cur-1] {
				pre = nums2[cur-1]
			}
			break
		}
		if second >= len(nums2) {
			cur := first + pos - i
			pre = result
			result = nums1[cur]
			if cur != 0 && pre < nums1[cur-1] {
				pre = nums1[cur-1]
			}
			break
		}
		if nums1[first] <= nums2[second] {
			pre = result
			result = nums1[first]
			first++
		} else {
			pre = result
			result = nums2[second]
			second++
		}
	}
	if ((len(nums1) + len(nums2)) % 2) != 0 {
		return float64(result)
	}
	return (float64(pre) + float64(result)) / 2

}

func main() {
	nums1 := []int{2}
	nums2 := []int{1, 3, 4}
	var result float64
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
