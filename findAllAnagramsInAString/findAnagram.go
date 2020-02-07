package findanagram

func findAnagrams(s string, p string) []int {
	result := []int{}
	left := 0
	right := len(p) - 1
	target := make([]byte, 26)
	match := make([]byte, 26)

	// 初始化target
	for i := 0; i < len(p); i++ {
		target[p[i]-'a'] += 1
		match[s[i]-'a'] += 1
	}
	for right < len(s) {
		if string(match) == string(target) {
			result = append(result, left)
		}
		match[s[left]-'a'] -= 1
		left++
		right++
		if right < len(s) {
			match[s[right]-'a'] += 1
		}
	}
	return result
}
