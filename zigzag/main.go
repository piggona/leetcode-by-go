package main

import "fmt"

func convert(s string, numRows int) string {
	// 边界条件
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	// 定义主要变量
	jump := 2*numRows - 2
	column := len(s) / jump
	if (len(s) % jump) != 0 {
		column++
	}
	var result []byte

	// 算法
	// 一般条件
	for i := 1; i <= numRows; i++ {
		pos := i - 1
		// add := numRows - i
		if i == 1 || i == numRows {
			for j := 0; j < column; j++ {
				if pos < len(s) {
					result = append(result, s[pos])
					pos = pos + jump
				}
			}
		}
		for j := 0; j < column; j++ {
			middle := 2*numRows*(j+1) - i - 2*j - 1
			if pos < len(s) && middle < len(s) {
				result = append(result, s[pos])
				result = append(result, s[middle])
				pos = pos + jump
			}
			if pos < len(s) && middle >= len(s) {
				result = append(result, s[pos])
				pos = pos + jump
			}
		}
	}
	return string(result)
}

func main() {
	s := "ABCDE"
	numRows := 4
	var result string
	result = convert(s, numRows)
	fmt.Println(result)
}
