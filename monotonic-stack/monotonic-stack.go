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

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	copy(ans, prices)
	st := []int{}
	for i, p := range prices {
		for len(st) >= 1 && p <= prices[st[len(st)-1]] {
			ans[st[len(st)-1]] = prices[st[len(st)-1]] - p
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
