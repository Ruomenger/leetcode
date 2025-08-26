package sliding_window

import "slices"

// 不定长滑动窗口

func lengthOfLongestSubstring(s string) int {
	ans := 0
	cnt := [128]bool{}
	for left, right := 0, 0; right < len(s); right++ {
		for cnt[s[right]] {
			cnt[s[left]] = false
			left++
		}
		cnt[s[right]] = true
		ans = max(ans, right-left+1)
	}
	return ans
}

func maximumLengthSubstring(s string) int {
	ans := 0
	cnt := [26]int{}
	for left, right := 0, 0; right < len(s); right++ {
		ch := s[right] - 'a'
		cnt[ch]++
		for cnt[ch] > 2 {
			cnt[s[left]-'a']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func longestSubarray(nums []int) int {
	ans := 0
	cnt := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		if cnt <= 1 {
			ans = max(ans, right-left)
		}
	}

	return ans
}

func minRemoval(nums []int, k int) int {
	slices.Sort(nums)
	ans := 1
	for left, right := 0, 0; right < len(nums); right++ {
		for nums[left]*k < nums[right] {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return len(nums) - ans
}

func equalSubstring(s string, t string, maxCost int) int {
	ans := 0
	length := len(s)
	cost := 0
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for left, right := 0, 0; right < length; right++ {
		cost += abs(int(s[right]) - int(t[right]))
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		ans = max(right-left+1, ans)
	}
	return ans
}

func totalFruit(fruits []int) int {
	fruitMap := make(map[int]int)
	ans := 0
	cnt := 0
	for left, right := 0, 0; right < len(fruits); right++ {
		fruitMap[fruits[right]]++
		cnt++
		for len(fruitMap) > 2 {
			fruitMap[fruits[left]]--
			cnt--
			if fruitMap[fruits[left]] == 0 {
				delete(fruitMap, fruits[left])
			}
			left++
		}
		ans = max(ans, cnt)
	}
	return ans
}

func maximumUniqueSubarray(nums []int) int {
	ans := 0
	numMap := make(map[int]int, 0)
	sum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		num := nums[right]
		sum += num
		numMap[num]++
		for numMap[num] > 1 {
			numMap[nums[left]]--
			sum -= nums[left]
			left++
			if numMap[nums[left]] == 0 {
				delete(numMap, nums[left])
			}
		}
		ans = max(ans, sum)
	}
	return ans
}

func maxSubarrayLength(nums []int, k int) int {
	ans := 0
	numMap := make(map[int]int)
	for left, right := 0, 0; right < len(nums); right++ {
		numMap[nums[right]]++
		for numMap[nums[right]] > k {
			numMap[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxAnswer(answerKey, k, 'T'), maxAnswer(answerKey, k, 'F'))
}

func maxAnswer(answerKey string, k int, ch byte) int {
	sum := 0
	ans := 0
	for left, right := 0, 0; right < len(answerKey); right++ {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k {
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func longestOnes(nums []int, k int) int {
	ans := 0
	cnt := 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func minOperations(nums []int, x int) int {
	target := -x
	for _, num := range nums {
		target += num
	}
	if target < 0 {
		return -1
	}
	ans := -1
	sum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target {
			ans = max(ans, right-left+1)
		}
	}
	if ans < 0 {
		return -1
	}
	return len(nums) - ans
}

func longestSemiRepetitiveSubstring(s string) int {
	ans := 0
	cnt := 0
	for left, right := 0, 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			cnt++
		}
		for cnt > 1 {
			left++
			if s[left] == s[left-1] {
				cnt--
			}
		}
		ans = max(ans, right-left+1)
	}

	return ans

}

func maximumBeauty(nums []int, k int) int {
	slices.Sort(nums)
	ans := 0
	for left, right := 0, 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
