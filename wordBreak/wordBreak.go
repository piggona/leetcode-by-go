package wordBreak

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	wordMap := make(map[string]bool)
	for _, str := range wordDict {
		wordMap[str] = true
	}
	return wordRecur(s, &wordMap)
}

func wordRecur(s string, candidates *map[string]bool) bool {
	dict := *candidates
	for i := 1; i < len(s); i++ {
		temp := s[:i]
		if dict[temp] == true {
			res := wordRecur(s[i:], candidates)
			if res == true {
				return true
			}
		}
	}
	temp := s[:len(s)]
	if dict[temp] == true {
		return true
	}
	return false
}

func wordBreakDP(s string, wordDict []string) bool {
	mem := make([]bool, len(s)+1)
	mem[0] = true

	dict := make(map[string]bool)
	for _, i := range wordDict {
		dict[i] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[j] && dict[s[j:i]] {
				mem[i] = true
				break
			}
		}
	}

	return mem[len(s)]
}
