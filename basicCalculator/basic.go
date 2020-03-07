package basicCalculator

func calculate(s string) int {
	pri := make(map[byte]int)
	maker := []byte{}
	var curNum int
	cal := []int{}
	pri['+'] = 1
	pri['-'] = 1
	pri['*'] = 2
	pri['/'] = 2
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			if curNum != 0 {
				cal[len(cal)-1] = curNum*10 + int(s[i]-'0')
			} else {
				cal = append(cal, int(s[i]-'0'))
			}
			curNum = cal[len(cal)-1]
		} else if s[i] != ' ' {
			if len(maker) != 0 {
				var temp byte = maker[len(maker)-1]
				for pri[temp] > pri[s[i]] {
					maker = maker[:len(maker)-1]
					// 操作
					n1 := cal[len(cal)-1]
					n2 := cal[len(cal)-2]
					switch s[i] {
					case '+':
						cal[len(cal)-2] = n1 + n2
					case '-':
						cal[len(cal)-2] = n2 - n1
					case '*':
						cal[len(cal)-2] = n2 * n1
					case '/':
						cal[len(cal)-2] = n2 / n1
					}
					cal = cal[:len(cal)-1]
					temp = maker[len(maker)-1]
				}
			}
			maker = append(maker, s[i])
			curNum = 0
		} else {
			curNum = 0
		}
	}
	for len(cal) != 1 {
		temp := maker[len(maker)-1]
		n1 := cal[len(cal)-1]
		n2 := cal[len(cal)-2]
		switch temp {
		case '+':
			cal[len(cal)-2] = n1 + n2
		case '-':
			cal[len(cal)-2] = n2 - n1
		case '*':
			cal[len(cal)-2] = n2 * n1
		case '/':
			cal[len(cal)-2] = n2 / n1
		}
		cal = cal[:len(cal)-1]
		maker = maker[:len(maker)-1]
	}
	return cal[0]
}
