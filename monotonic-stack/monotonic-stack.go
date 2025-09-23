package monotonicstack

func dailyTemperatures(temperatures []int) []int {
	st := []int{}
	ans := make([]int, len(temperatures))
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
