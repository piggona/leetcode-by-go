package main

import (
	"fmt"
)

func grayCode(n int) []int {
	size := pow(n) - 1
	numMap := make(map[int]bool)
	result := []int{0}
	numMap[0] = true
	is, result := getCode(n, size, numMap, result)
	if is {
		return result
	}
	return nil
}

func binaryToGray(n int) int {
	return (n >> 1) ^ n
}

func grayToBinary(num int) int {
	var mask int
	for mask = num >> 1; mask != 0; mask = mask >> 1 {
		num = num ^ mask
	}
	return num
}

func getCode(n, size int, numMap map[int]bool, result []int) (bool, []int) {
	if size == 0 {
		return true, result
	}
	for i := 0; i < n; i++ {
		temp := result[len(result)-1]
		divide := getDivide(temp, n)
		pos := divide[i]
		new := 0
		if pos == 1 {
			new = temp - pow(i)
		} else {
			new = temp + pow(i)
		}

		if !numMap[new] {
			list := append(result, new)
			newMap := numMap
			newMap[new] = true
			is, res := getCode(n, size-1, newMap, list)
			if is {
				return true, res
			}
		}
	}
	return false, nil
}

func getDivide(num, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		temp := num % 2
		result = append(result, temp)
		num = num / 2
	}
	return result
}

func pow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 2
	}
	return result
}

func main() {
	input := 3
	fmt.Println(grayCode(input))
}
