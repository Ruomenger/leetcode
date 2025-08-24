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
