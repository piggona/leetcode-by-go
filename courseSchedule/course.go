package courseSchedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	bitmap := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		bitmap[i] = make([]bool, numCourses+1)
	}

	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		pre := prerequisites[i][1]
		// 空的
		loop := noLoop(bitmap, course, pre)
		if !loop {
			return false
		} else {
			bitmap[course][pre] = true
			bitmap[course][numCourses] = true
		}
	}
	return true
}

func noLoop(bitmap [][]bool, course, pre int) bool {
	if bitmap[pre][course] {
		return false
	}
	if !bitmap[pre][len(bitmap[0])-1] {
		return true
	}
	temp := bitmap[pre]
	for i := 0; i < len(bitmap[0])-1; i++ {
		if temp[i] {
			res := noLoop(bitmap, course, i)
			if !res {
				return false
			}
		}
	}
	return true
}
