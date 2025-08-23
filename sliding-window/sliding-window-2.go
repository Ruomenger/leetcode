package sliding_window

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
