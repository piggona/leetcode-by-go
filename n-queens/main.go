package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	result := [][]string{}
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
	for i := 0; i < len(candidate); i++ {
		tt := make([]int, len(temp))
		tempCandiates := make([]int, len(candidate))
		copy(tempCandiates, candidate)
		copy(tt, temp)
		if len(tt) == 0 || abs(candidate[i]-tt[len(tt)-1]) > 1 {
			tt = append(tt, candidate[i])
			tempCandiates[i], tempCandiates[0] = tempCandiates[0], tempCandiates[i]
			genNums(tempCandiates[1:], tt, result)
		} else {
			continue
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
	return result
}

func main() {
	num := 4
	fmt.Println(notAdjust(num))
}
