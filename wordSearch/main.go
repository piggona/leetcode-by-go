package main

import "fmt"

func exist(board [][]byte, word string) bool {

	// 找到存在的word的第一个字符
	// 之后开始深度遍历
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				board[i][j] = '#'
				if depth(board, i, j, 1, word) {
					return true
				}
				board[i][j] = word[0]
			}
		}
	}
	return false
}

func depth(board [][]byte, i, j, k int, word string) bool {
	// i横轴坐标，j纵轴坐标，k当前匹配到word的位置
	// 以i，j为起点，向四周递归。递归返回后
	if k >= len(word) {
		return true
	}
	result := false
	// 向上
	temp := i - 1
	if temp >= 0 && board[temp][j] == word[k] {
		board[temp][j] = '#'
		result = result || depth(board, temp, j, k+1, word)
		board[temp][j] = word[k]
	}
	// 向下
	temp = i + 1
	if temp < len(board) && board[temp][j] == word[k] {
		board[temp][j] = '#'
		result = result || depth(board, temp, j, k+1, word)
		board[temp][j] = word[k]
	}
	// 向左
	temp = j - 1
	if temp >= 0 && board[i][temp] == word[k] {
		board[i][temp] = '#'
		result = result || depth(board, i, temp, k+1, word)
		board[i][temp] = word[k]
	}
	// 向右
	temp = j + 1
	if temp < len(board[0]) && board[i][temp] == word[k] {
		board[i][temp] = '#'
		result = result || depth(board, i, temp, k+1, word)
		board[i][temp] = word[k]
	}
	return result
}

func main() {
	word := "ABCB"
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, word))
}
