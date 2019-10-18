package main

import (
	"fmt"
)

func getFactorial(n int) int {
	result := 1
	for n != 1 {
		result *= n
		n--
	}
	return result
}

func getPermutation(n int, k int) string {
	factorial := 1
	numbers := []int{1}
	for i := 2; i <= n; i++ {
		factorial *= i
		numbers = append(numbers, i)
	}

	ret := make([]byte, 0)
	for i := n; i > 0; i-- {
		factorial = factorial / i
		pos := (k - 1) / factorial
		k = k - factorial*pos
		num := 0
		num, numbers = getN(numbers, pos)
		ret = append(ret, byte(num)+'0')
	}
	return string(ret)
}

// getN
func getN(nums []int, n int) (int, []int) {
	result := nums[n]
	nums = append(nums[:n], nums[n+1:]...)
	return result, nums
}

func main() {
	input := 1
	k := 1
	fmt.Println(getPermutation(input, k))
}
