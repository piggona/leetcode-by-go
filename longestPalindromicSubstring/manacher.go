package main

import "fmt"

func longestPalindrome(s string) string {
	// 算法数据声明
	var exStr []byte
	var result string
	j := 0
	pivot := 0
	maxLength := 0
	// 边界条件
	if len(s) == 1 {
		return s
	}
	//算法
	// 初始化算法数据结构
	for _, str := range []byte(s) {
		exStr = append(exStr, byte('#'), str)
	}
	exStr = append(exStr, '#')
	var RL [2001]int
	for i := 0; i < len(exStr); i++ {
		left := i
		right := i
		j = i - pivot - 1
		if i < maxLength {
			move := min((pivot - j), (RL[j] / 2), maxLength-i)
			left = left - move
			right = right + move
		}
		for right != len(exStr)-1 && left != 0 {
			if exStr[left-1] == exStr[right+1] {
				left--
				right++
			} else {
				break
			}
		}
		if right > maxLength {
			maxLength = right
			pivot = i
		}
		RL[i] = right - left + 1
	}
	pivot = max(RL[0:])
	maxLength = RL[pivot] / 2
	left := pivot - maxLength
	right := pivot + maxLength
	result = getMaxSub(exStr[left:right])
	return result
}
func getMaxSub(str []byte) string {
	result := []byte{}
	for _, val := range str {
		if val != byte('#') {
			result = append(result, val)
		}
	}
	return string(result)
}
func max(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[result] < nums[i] {
			result = i
		}
	}
	return result
}

func min(i1, i2, i3 int) int {
	result := i1
	if i1 > i2 {
		result = i2
	}
	if i3 < result {
		result = i3
	}
	return result
}

func main() {
	s := "bba"
	result := longestPalindrome(s)
	fmt.Println(result)
}
