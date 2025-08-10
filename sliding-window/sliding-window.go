package sliding_window

import (
	"math"
	"slices"
)

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

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxSum := math.MinInt
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		if right-left+1 < k {
			continue
		}
		maxSum = max(maxSum, sum)
		sum -= nums[left]
		left++
	}
	return float64(maxSum) / float64(k)
}
