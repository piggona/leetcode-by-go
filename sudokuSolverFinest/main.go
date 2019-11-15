package main

import "fmt"

var choices = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudoku(board [][]byte) {
	choiceMaker(board, 0, 0)
}

func choiceMaker(board [][]byte, r, c int) bool {
	if r == 9 {
		return true
	}
	// next spot
	nR, nC := r, c+1 // 下一个位置
	if nC == 9 {     // 跳到下一行
		nR, nC = r+1, 0
	}
	// fixed
	if board[r][c] != '.' {
		return choiceMaker(board, nR, nC) // 如果已经填好数字，跳到下一个
	}
	// try all possibilities!
	for _, option := range choices {
		if validChoice(board, option, r, c) {
			// place choice
			board[r][c] = option
			if !choiceMaker(board, nR, nC) {
				// remove
				board[r][c] = '.'
			} else {
				return true
			}
		}
	}
	return false
}

func validChoice(board [][]byte, choice byte, r, c int) bool {
	// row & col
	for i := 0; i < 9; i++ {
		if board[r][i] == choice || board[i][c] == choice {
			return false
		}
	}
	// square 确定方块内没有与Choice相同的数字
	rowLock, colLock := 0, 0
	if 2 < c && c < 6 { // mid
		colLock = 3
	} else if c > 5 { // right
		colLock = 6
	}
	if 2 < r && r < 6 { // mid
		rowLock = 3
	} else if r > 5 { // bot
		rowLock = 6
	}
	for i := 0; i < 3; i++ {
		if board[rowLock+i][colLock] == choice ||
			board[rowLock+i][colLock+1] == choice ||
			board[rowLock+i][colLock+2] == choice {
			return false
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
	// temp := []int{1, 2, 3, 4}
	// count := stepMult(len(temp))
	// for i := 0; i < count; i++ {
	// 	fmt.Println(nextTemp(temp))
	// }
	// bMap := []byte{'1', '2', '4', '7'}
	// clist := []map[byte]bool{
	// 	{'1': true, '2': true, '4': true, '6': true, '9': true},
	// 	{'2': true, '3': true, '4': true, '7': true, '8': true},
	// 	{'2': true, '3': true, '4': true, '7': true},
	// 	{'2': true},
	// }
	// fmt.Println(matchList(bMap, clist))
}
