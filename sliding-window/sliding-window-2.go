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
