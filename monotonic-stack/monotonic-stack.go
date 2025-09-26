package monotonicstack

import "slices"

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

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := []int{}
	for i := 0; i < 2*n; i++ {
		num := nums[i%n]
		for len(st) > 0 && num > nums[st[len(st)-1]] {
			ans[st[len(st)-1]] = num
			st = st[:len(st)-1]
		}
		if i < n {
			st = append(st, i)
		}
	}
	return ans
}

func carFleet(target int, position []int, speed []int) int {
	type pair struct {
		pos  int
		hour float64
	}
	cars := make([]pair, len(position))
	for i, pos := range position {
		cars = append(cars, pair{
			pos:  pos,
			hour: float64(target-pos) / float64(speed[i]),
		})
	}
	slices.SortFunc(cars, func(a, b pair) int { return a.pos - b.pos })
	st := []pair{}
	for _, car := range cars {
		for len(st) > 0 && car.hour >= st[len(st)-1].hour {
			st = st[:len(st)-1]
		}
		st = append(st, car)
	}
	return len(st)
}

func largestRectangleArea(heights []int) int {
	ans := 0
	left := make([]int, len(heights))
	st := []int{-1}
	for i, h := range heights {
		for len(st) > 1 && heights[st[len(st)-1]] >= h {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	right := make([]int, len(heights))
	st = st[:1]
	st[0] = len(heights)
	for i, h := range slices.Backward(heights) {
		for len(st) > 1 && heights[st[len(st)-1]] >= h {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	for i, h := range heights {
		ans = max(ans, h*(right[i]-left[i]-1))
	}

	return ans
}
