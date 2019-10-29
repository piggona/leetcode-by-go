package main

import "fmt"

func simplifyPath(path string) string {
	pos := 1
	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	for pos < len(path) {
		switch path[pos] {
		case '/':
			if path[pos-1] == '/' {
				path = path[:pos] + path[pos+1:]
			} else if path[pos-1] == '.' {
				if path[pos-2] == '.' {
					temp := pos - 4
					if temp < 0 {
						path = path[:1] + path[pos+1:]
						pos = 1
					} else {
						path = path[:temp] + path[pos+1:]
						pos -= 4
					}
				} else {
					path = path[:pos-1] + path[pos+1:]
					pos--
				}
			} else {
				pos++
			}
			break
		default:
			pos++
			break
		}
	}
	if len(path) > 1 {
		return path[:len(path)-1]
	}
	return path
}

func main() {
	input := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(input))
}
