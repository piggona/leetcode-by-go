package main

import (
	"fmt"
	"sort"
	"strings"
)

func getPermutation(n int, k int) string {
	if n == 1 {
		return string([]byte{'1'})
	}
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	for i := 1; i < k; i++ {
		nextPermutation(nums)
	}
	result := strings.Builder{}
	for i := 0; i < n; i++ {
		result.WriteByte(byte(nums[i]) + '0')
	}
	return result.String()
}

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			right := len(nums) - 1
			for nums[i] > nums[right] {
				right--
			}
			nums[i], nums[right] = nums[right], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
	return
}

func main() {
	input := 4
	k := 9
	fmt.Println(getPermutation(input, k))
}
