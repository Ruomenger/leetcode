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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	num2Idx := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		num2Idx[nums1[i]] = i
	}
	ans := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = -1
	}
	stack := make([]int, 0)
	for _, num := range nums2 {
		for len(stack) != 0 && num > stack[len(stack)-1] {
			ans[num2Idx[stack[len(stack)-1]]] = num
			stack = stack[:len(stack)-1]
		}
		if _, ok := num2Idx[num]; ok {
			stack = append(stack, num)
		}
	}
	return ans
}
