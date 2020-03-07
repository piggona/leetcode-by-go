package isNumber

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	var (
		cur      = 1
		hasPoint bool
		left     = 0
		right    = len(s) - 1
	)
	for left < len(s) && s[left] == ' ' {
		left++
	}
	for right >= 0 && s[right] == ' ' {
		right--
	}
	if left > right {
		return false
	}
	s = s[left : right+1]
	if s[0] != '+' && s[0] != '-' && s[0] != '.' && (s[0] < '0' || s[0] > '9') {
		return false
	}
	if (s[0] == '+' || s[0] == '-' || s[0] == '.') && cur >= len(s) {
		return false
	}
	if s[0] == '.' {
		hasPoint = true
	}
	if s[0] == '.' && (s[1] == 'e' || s[1] == 'E') {
		return false
	}
	for cur < len(s) && (s[cur] != 'e' && s[cur] != 'E') {
		temp := s[cur]
		if temp == '.' && ((cur+1 == len(s)) || (s[cur+1] >= '0' && s[cur+1] <= '9')) && !hasPoint {
			hasPoint = true
		} else if temp >= '0' && temp <= '9' {
		} else {
			return false
		}
		cur++
	}
	if cur == len(s) {
		return true
	}
	return AfterE(s[cur+1:])
}

func AfterE(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] != '+' && s[0] != '-' && (s[0] < '0' || s[0] > '9') {
		return false
	}
	if (s[0] == '+' || s[0] == '-') && len(s) == 1 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
