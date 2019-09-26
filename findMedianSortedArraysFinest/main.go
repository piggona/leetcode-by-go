package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var result float64
	// 数据前提条件
	length := len(nums1) + len(nums2)
	odd := (length % 2) != 0
	var symmetry int
	if odd {
		symmetry = length / 2
	} else {
		symmetry = length/2 - 1
	}
	// 算法数据
	remain := symmetry

	// 算法
	// 1. 边界情况
	if len(nums2) == 0 {
		if odd {
			pos := length / 2
			return float64(nums1[pos])
		}
		pos := length / 2
		return (float64(nums1[pos]) + float64(nums1[pos-1])) / 2
	}
	if length == 2 {
		return (float64(nums1[0]) + float64(nums2[0])) / 2
	}
	// 主算法
	result = recursiveMedian(nums1, nums2, remain, odd)
	return result
}

func recursiveMedian(nums1 []int, nums2 []int, remain int, odd bool) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 要考虑边界条件优先级
	if remain == 0 {
		if odd {
			if nums1[0] < nums2[0] {
				return float64(nums1[0])
			} else {
				return float64(nums2[0])
			}
		}
		result1, result2 := get2Small(nums1, nums2)
		return (float64(result1) + float64(result2)) / 2
	}
	if len(nums2) == 1 {
		remain--
		if nums1[0] < nums2[0] {
			nums1 = nums1[1:]
			return recursiveMedian(nums1, nums2, remain, odd)
		} else {
			pos := remain
			if odd {
				return float64(nums1[pos])
			}
			return (float64(nums1[pos]) + float64(nums1[pos+1])) / 2
		}
	}
	if remain == 1 {
		remain--
		if nums1[0] < nums2[0] {
			nums1 = nums1[1:]
		} else {
			nums2 = nums2[1:]
		}
		return recursiveMedian(nums1, nums2, remain, odd)
	}
	// 一般情况
	x := len(nums2)
	pos := (2*x-1)/2 - 1
	posRe := remain/2 - 1
	if pos > posRe {
		pos = posRe
	}
	remain = remain - pos - 1
	if nums1[pos] < nums2[pos] {
		nums1 = nums1[(pos + 1):]
	} else if nums1[pos] > nums2[pos] {
		nums2 = nums2[(pos + 1):]
	} else if nums1[pos] == nums2[pos] {
		remain = remain - pos - 1
		nums1 = nums1[(pos + 1):]
		nums2 = nums2[(pos + 1):]
	}
	return recursiveMedian(nums1, nums2, remain, odd)

}

// type SearchList interface {
// 	Len() int
// }

// type judgeSearch

// func MidSearch()

func larger(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func smaller(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func get2Small(nums1 []int, nums2 []int) (int, int) {
	var result [2]int
	if (len(nums1) + len(nums2)) == 2 {
		return nums1[0], nums2[0]
	}
	i := 0
	j := 0
	for k := 0; k < 2; k++ {
		if j == 1 && len(nums2) == 1 {
			result[k] = nums1[i]
			break
		}
		if nums1[i] < nums2[j] {
			result[k] = nums1[i]
			i++
		} else {
			result[k] = nums2[j]
			j++
		}
	}
	return result[0], result[1]
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6, 7, 8, 9}
	var result float64
	result = findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}
