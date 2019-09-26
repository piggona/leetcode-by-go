package main

import "fmt"

// func isValidSudoku(board [][]byte) bool {
// 	enum := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
// 	lineMap := [9]map[byte]bool{}
// 	columnMap := [9]map[byte]bool{}
// 	// 行与列的候选
// 	for i := 0; i < 9; i++ {
// 		line := make(map[byte]bool)
// 		column := make(map[byte]bool)
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != '.' {
// 				line[board[i][j]] = true
// 			}
// 			if board[j][i] != '.' {
// 				column[board[j][i]] = true
// 			}
// 		}
// 		lineMap[i] = line
// 		columnMap[i] = column
// 	}
// 	// 合并行与列的候选
// 	candidate := [9][9]map[byte]bool{}
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			temp := make(map[byte]bool)
// 			line := lineMap[i]
// 			column := columnMap[j]
// 			for _, num := range enum {
// 				if !(line[num] || column[num]) {
// 					temp[num] = true
// 				}
// 			}
// 			candidate[i][j] = temp
// 		}
// 	}

// 	// 3 x 3小方格试验
// 	// 在每个方格中，先排除已有的byte，查看每个空元素，是否满足剩下的需求
// 	result := true
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			bMap := make(map[byte]bool)
// 			r := i * 3
// 			c := j * 3
// 			place := make(map[int][]int)
// 			for m := r; m < 3; m++ {
// 				for n := c; n < 3; n++ {
// 					if board[m][n] != '.' {
// 						bMap[board[m][n]] = true
// 					} else {
// 						place[m] = append(place[m], n)
// 					}
// 				}
// 			}
// 			clist := []map[byte]bool{}
// 			for key, value := range place {
// 				for _, n := range value {
// 					temp := candidate[key][n]
// 					clist = append(clist, temp)
// 				}
// 			}
// 			bTemp := []byte{}
// 			for _, num := range enum {
// 				if bMap[num] == false {
// 					bTemp = append(bTemp, num)
// 				}
// 			}
// 			result = result && matchList(bTemp, clist)
// 		}
// 	}

// 	return result
// }

func isValidSudoku(board [][]byte) bool {
	var bitmaps [3][9][9]int
	// 第一维表示方向（即行、列、方块内），第二维表示方向上的位置，第三维表示方向上某位置上的数字
	res := true
	for i, line := range board {
		for j, cell := range line {
			if cell == '.' {
				continue
			}
			if bitmaps[0][i][int(cell-'1')] == 0 &&
				bitmaps[1][j][int(cell-'1')] == 0 &&
				bitmaps[2][(i/3)*3+j/3][int(cell-'1')] == 0 {
				bitmaps[0][i][int(cell-'1')], bitmaps[1][j][int(cell-'1')],
					bitmaps[2][(i/3)*3+j/3][int(cell-'1')] = 1, 1, 1
			} else {
				res = false
				break
			}
		}
	}
	return res
}

func matchList(bMap []byte, clist []map[byte]bool) bool {
	result := false
	for i, num := range bMap {
		for id, list := range clist {
			if list[num] == true {
				if len(bMap) == 1 {
					return true
				}

				temp := make([]byte, len(bMap))
				tlist := make([]map[byte]bool, len(clist))
				copy(temp, bMap)
				copy(tlist, clist)

				temp[i], temp[0] = temp[0], temp[i]
				temp = temp[1:]
				tlist[0], tlist[id] = tlist[id], tlist[0]
				tlist = tlist[1:]
				result = result || matchList(temp, tlist)
			}
		}
	}
	return result
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
	fmt.Println(isValidSudoku(board))
	// bMap := []byte{'1', '2', '4', '7'}
	// clist := []map[byte]bool{
	// 	{'1': true, '2': true, '4': true, '6': true, '9': true},
	// 	{'2': true, '3': true, '4': true, '7': true, '8': true},
	// 	{'2': true, '3': true, '4': true, '7': true},
	// 	{'2': true},
	// }
	// fmt.Println(matchList(bMap, clist))
}
