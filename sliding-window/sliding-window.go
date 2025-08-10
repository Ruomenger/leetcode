package sliding_window

import "slices"

// maxVowels 1456. 定长子串中元音的最大数目
func maxVowels(s string, k int) int {
	ans := 0
	cnt := 0
	for left, right := 0, 0; right < len(s); right++ {
		if slices.Contains([]byte{'a', 'e', 'i', 'o', 'u'}, s[right]) {
			cnt++
		}
		if right-left+1 < k {
			continue
		}
		ans = max(ans, cnt)
		if slices.Contains([]byte{'a', 'e', 'i', 'o', 'u'}, s[left]) {
			cnt--
		}
		left++
	}

	return ans
}
