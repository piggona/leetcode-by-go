package reconstruction

import "sort"

type pep [][]int

func (p pep) Len() int {
	return len(p)
}

func (p pep) Less(i, j int) bool {
	if p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1]) {
		return true
	}
	return false
}

func (p pep) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func ReconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		res[i] = make([]int, 2)
	}
	sort.Sort(pep(people))
	for i := 0; i < len(people); i++ {
		temp := people[i]
		if temp[1] == 0 {
			count := temp[1]
			pos := 0
			for count >= 0 {
				if res[pos][0] == 0 || res[pos][0] == temp[0] {
					count--
				}
				pos++
			}
			pos--
			for res[pos][0] != 0 {
				pos++
			}
			res[pos] = temp
		} else {
			count := temp[1]
			pos := 0
			for count > 0 {
				if res[pos][0] == 0 || res[pos][0] == temp[0] {
					count--
				}
				pos++
			}
			for res[pos][0] != 0 {
				pos++
			}
			res[pos] = temp
		}
	}
	return res
}
