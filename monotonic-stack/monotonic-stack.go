package monotonicstack

import (
	"slices"
)

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

func maximumScore(nums []int, k int) int {
	ans := 0
	left := make([]int, len(nums))
	st := []int{-1}
	for i, h := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] >= h {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	right := make([]int, len(nums))
	st = st[:1]
	st[0] = len(nums)
	for i, h := range slices.Backward(nums) {
		for len(st) > 1 && nums[st[len(st)-1]] >= h {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	for i, h := range nums {
		l, r := left[i], right[i]
		if l < k && k < r {
			ans = max(ans, h*(right[i]-left[i]-1))
		}
	}

	return ans
}

func maximalRectangle(matrix [][]byte) (ans int) {
	heights := make([]int, len(matrix[0]))
	for _, row := range matrix {
		// 计算底边为 row 的柱子高度
		for j, c := range row {
			if c == '0' {
				heights[j] = 0 // 柱子高度为 0
			} else {
				heights[j]++ // 柱子高度加一
			}
		}
		ans = max(ans, largestRectangleArea(heights)) // 调用 84 题代码
	}
	return
}

func trap(height []int) int {
	ans := 0
	n := len(height)
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := range n {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func trap2(height []int) int {
	ans := 0
	n := len(height)
	leftMax, rightMax := 0, 0
	left, right := 0, n-1
	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func trap3(height []int) int {
	ans := 0
	st := []int{}
	for i, h := range height {
		for len(st) > 0 && height[st[len(st)-1]] <= h {
			bottomH := height[st[len(st)-1]]
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			dh := min(height[left], h) - bottomH
			ans += dh * (i - left - 1)
		}
		st = append(st, i)
	}
	return ans
}

func sumSubarrayMins(arr []int) int {
	st := []int{-1}
	left := make([]int, len(arr))
	for i, num := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= num {
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	st = []int{len(arr)}
	right := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for len(st) > 1 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1]
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}
	ans := 0
	for i, num := range arr {
		ans += num * (i - left[i]) * (right[i] - i)
	}
	return ans % (1e9 + 7)
}

func maxSums(nums []int) int64 {
	ans := int64(0)
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, num := range nums {
		for len(st) > 1 && nums[st[len(st)-1]] <= num {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	for i, num := range nums {
		ans += int64(i-left[i]) * int64(right[i]-i) * int64(num)
	}
	return ans
}

func subArrayRanges(nums []int) int64 {
	ans := maxSums(nums)
	for i := range nums {
		nums[i] *= -1
	}
	return ans + maxSums(nums)
}

func maxSumMinProduct(nums []int) int {
	ans := 0
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, num := range nums {
		for len(st) > 1 && num <= nums[st[len(st)-1]] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	for i, num := range nums {
		ans = max(ans, num*(prefixSum[right[i]]-prefixSum[left[i]+1]))
	}

	return ans % (1e9 + 7)
}
