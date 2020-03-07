package strtoint

import (
	"math"
	"strconv"
)

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err == nil {
		if i < math.MinInt32 {
			return math.MinInt32
		}
		if i > math.MaxInt32 {
			return math.MaxInt32
		}
		return i
	}
	flag := false
	start := 0
	end := 0
	var id int
	var s rune
	for id = 0; id < len(str); id++ {
		s = rune(str[id])
		if !flag && s == ' ' {
			continue
		}
		if !flag && (s < '0' || s > '9') {
			if s == '-' || s == '+' {
				start = id
				flag = true
				continue
			} else {
				return 0
			}
		}
		if !flag && s >= '0' && s <= '9' {
			start = id
			flag = true
			continue
		}
		if flag && (s < '0' || s > '9' || s == ' ') {
			end = id
			break
		}
	}
	if flag && id == len(str) {
		end = len(str)
	}
	op := 1
	num := 0
	for i := start; i < end; i++ {
		if str[i] == '-' {
			op = -1
		} else if str[i] == '+' {
			continue
		} else {
			temp := int(str[i] - '0')
			num = num*10 + temp
			if num*op < math.MinInt32 {
				return math.MinInt32
			}
			if num*op > math.MaxInt32 {
				return math.MaxInt32
			}
		}
	}
	num = op * num
	if num < math.MinInt32 {
		return math.MinInt32
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return num
}
