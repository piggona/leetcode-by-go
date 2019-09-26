package main

import "fmt"

func isMatchFalse(s string, p string) bool {
	si := 0
	pi := 0
	var pre byte
	// var val byte
	// flag := false

	// 初始条件
	if p == ".*" {
		return true
	}
	// 以.*为界分离p字符串
	strMap := map[int]string{}
	max := 0
	left := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '*' && pre == '.' {
			strMap[max] = p[left : i-1]
			max++
			left = i + 1
		}
		pre = p[i]
	}
	strMap[max] = p[left:len(p)]
	// 算法
	// 先匹配第一个子串
	p0 := strMap[0]
	for pi = 0; pi < len(p0); {
		if si >= len(s) { // p要匹配的东西要多于s
			if max == 0 {
				if pi == len(p0)-1 && p0[pi] == '*' {
					return true
				}
				if pi == len(p0)-2 && p0[pi+1] == '*' {
					return true
				}
			}
			return false
		}
		if s[si] == p0[pi] || p0[pi] == '.' {
			pre = p0[pi]
			si++
			pi++
		} else if p0[pi] == '*' {
			if s[si] == pre {
				si++
			} else {
				pi++
			}
		} else {
			if pi+1 < len(p0) && p0[pi+1] == '*' {
				pi = pi + 2
			} else {
				return false
			}
		}
	}
	if max == 0 && s[si:] == "" {
		return true
	}
	// 判断之后的子串，先判断strMap[i]是否存在，之后再判断有没有符合表达式的串
	for i := 1; i <= max; i++ {
		temp, ok := strMap[i]
		if ok {
			if temp == "" && i == max {
				return true
			}
			for pi = 0; pi < len(temp); {
				if si >= len(s) {
					if i == max {
						if pi == len(temp)-1 && temp[pi] == '*' {
							return true
						}
						if pi == len(temp)-2 && temp[pi+1] == '*' {
							return true
						}
					}
					return false
				}
				if s[si] == temp[pi] || temp[pi] == '.' {
					pre = temp[pi]
					si++
					pi++
				} else if temp[pi] == '*' {
					if s[si] == pre {
						si++
					} else {
						pi++
					}
				} else {
					if pi+1 < len(temp) && temp[pi+1] == '*' {
						pi = pi + 2
					} else {
						si++
					}
				}
			}
		}
	}
	if si < len(s) {
		return false
	}
	return true
}

func isMatchFlase2(s string, p string) bool {
	left := 0
	right := 0
	pi := 0
	var last byte
	strMap := map[int]string{}
	max := 0
	// 初始条件
	if p == ".*" {
		return true
	}
	// 以*为界分离p字符串
	strMap, max = getSplitByStar(p)
	for i := 0; i <= max; i++ {
		temp := strMap[i]
		if temp == "" && i == max {
			// 余下还有*
			for right < len(s) {
				if s[right] != last {
					break
				}
				right++
			}
			if right == len(s) {
				return true
			} else {
				return false
			}
		}
		if i == max {
			for len(s)-right > len(temp) {
				if last == s[right] || last == '.' {
					left++
					right = left
				} else {
					return false
				}
			}
		}
		for pi = 0; pi < len(temp); pi = pi {
			if right < len(s) && s[right] != temp[pi] {
				if last == s[right] || last == '.' {
					left++
					right = left
					pi = 0
					continue
				} else if temp[pi] == '.' {
				} else {
					break
				}
			}
			right++
			pi++
		}
		if pi == len(temp) { // 匹配成功
			last = temp[len(temp)-1]
			left = right
			right = left
		} else {
			if i >= max || pi != len(temp)-1 {
				return false
			}
		}
	}
	if left != len(s) {
		return false
	}
	return true
}

func getSplitByStar(p string) (map[int]string, int) {
	strMap := map[int]string{}
	max := 0
	left := 0
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			strMap[max] = p[left:i]
			max++
			left = i + 1
		}
	}
	strMap[max] = p[left:len(p)]
	return strMap, max
}

// func isMatch(s string, p string) bool {
// 	strP := []string{}
// 	strP = getSplitByte(p)
// 	dp := make([][]bool)
// 	return false
// }

func matchByte(s byte, p string) bool {
	return s == p[0] || p[0] == '.'
}

func getSplit(p string) []string {
	result := []string{}
	for i := 1; i < len(p); i++ {
		if p[i] == '*' {
			if i-2 >= 0 {
				result = append(result, p[:i-1])
			}
			if i+1 < len(p) {
				result = append(result, p[i-1:i+1])
				p = p[i+1:]
			} else {
				result = append(result, p[i-1:])
				return result
			}
			i = 1
		}
	}
	result = append(result, p)
	return result
}

func getSplitByte(p string) []string {
	result := []string{}
	for i := 0; i < len(p); i++ {
		if i+1 >= len(p) {
			result = append(result, string(p[i]))
		} else {
			if p[i+1] == '*' {
				if i+2 >= len(p) {
					result = append(result, p[i:])
				} else {
					result = append(result, p[i:i+2])
				}
				i = i + 1
			} else {
				result = append(result, string(p[i]))
			}
		}
	}
	return result
}

func main() {
	// s := "mississippi"
	// p := "mis*is*ip*."
	// s := "aaabdzabc"
	// p := "c*a*.*bc"
	// s := "baaaabaaab"
	// p := "b.*aba*ab"
	// s := "aaaba"
	p := "asdfb*.*c*a"
	fmt.Println(getSplit(p))
	// fmt.Println(isMatch(s, p))
}
