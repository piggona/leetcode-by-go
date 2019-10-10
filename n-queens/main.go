package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	result := [][]string{}
	for _, temp := range notAdjust(n) {
		result = append(result, displayQueens(temp))
	}
	return result
}

// noAdjust 生成不连续数列
func notAdjust(n int) [][]int {
	result := [][]int{}
	candidate := make([]int, n)
	for i := 0; i < n; i++ {
		candidate[i] = i
	}
	temp := []int{}
	genNums(candidate, temp, &result)
	return result
}

func genNums(candidate []int, temp []int, result *[][]int) {
	if len(candidate) == 0 {
		*result = append(*result, temp)
		return
	}
	for id, num := range candidate {
		flag := false
		for i, t := range temp {
			if abs(i-len(temp)) == abs(t-num) {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		if len(temp) == 0 || abs(temp[len(temp)-1]-num) > 1 {
			ct := append([]int{}, candidate...)
			tt := append([]int{}, temp...)
			genNums(append(ct[:id], ct[id+1:]...), append(tt, num), result)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// displayQueens 对单个棋盘（不连续数列）合成字符串数组输出
func displayQueens(nums []int) []string {
	result := []string{}
	for _, num := range nums {
		temp := strings.Builder{}
		for i := 0; i < len(nums); i++ {
			if i == num {
				temp.WriteByte('Q')
			} else {
				temp.WriteByte('.')
			}
		}
		result = append(result, temp.String())
	}
	return result
}

func main() {
	num := 5
	fmt.Println(notAdjust(num))
}
