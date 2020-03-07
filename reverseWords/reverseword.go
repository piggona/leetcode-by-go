package reverseword

import "strings"

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	str := strings.Split(s, " ")
	res := make([]string, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		if len(str[i]) != 0 {
			res = append(res, str[i])
		}
	}
	return strings.Join(res, " ")
}
