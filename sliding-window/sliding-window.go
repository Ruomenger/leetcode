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

func numOfSubarrays(arr []int, k int, threshold int) int {
	ans := 0
	maxValue := k * threshold
	sum := 0
	for left, right := 0, 0; right < len(arr); right++ {
		sum += arr[right]
		if right-left+1 < k {
			continue
		}
		if sum >= maxValue {
			ans++
		}
		sum -= arr[left]
		left++
	}
	return ans
}

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = -1
	}
	sum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		if right-left+1 < 2*k+1 {
			continue
		}
		ans[right-k] = sum / (2*k + 1)
		sum -= nums[left]
		left++
	}
	return ans
}

func maxSum(nums []int, m int, k int) int64 {
	numsMap := make(map[int]int)
	ans := 0
	sum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		numsMap[nums[right]]++
		sum += nums[right]
		if right-left+1 < k {
			continue
		}
		cnt := len(numsMap)
		if cnt >= m {
			ans = max(ans, sum)
		}
		numsMap[nums[left]]--
		if numsMap[nums[left]] == 0 {
			delete(numsMap, nums[left])
		}
		sum -= nums[left]
		left++
	}
	return int64(ans)
}
