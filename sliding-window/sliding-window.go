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

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	m := n - k
	sum := 0
	totalSum := 0
	ans := math.MaxInt
	for left, right := 0, 0; right < len(cardPoints); right++ {
		totalSum += cardPoints[right]
		sum += cardPoints[right]
		if right-left+1 < m {
			continue
		}
		ans = min(sum, ans)
		sum -= cardPoints[left]
		left++
	}
	if n == k {
		ans = 0
	}
	return totalSum - ans
}

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	ans := 0
	sum := 0
	sum0 := 0
	for left, right := 0, 0; right < len(customers); right++ {
		if grumpy[right] == 0 {
			sum0 += customers[right]
		} else {
			sum += customers[right]
		}
		if right-left+1 < minutes {
			continue
		}
		ans = max(ans, sum)
		if grumpy[left] == 1 {
			sum -= customers[left]
		}
		left++
	}
	return ans + sum0
}

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	freeTime := make([]int, n+1)
	freeTime[0] = startTime[0]
	for i := 1; i < n; i++ {
		freeTime[i] = startTime[i] - endTime[i-1]
	}
	freeTime[n] = eventTime - endTime[n-1]
	sum := 0
	ans := 0
	for left, right := 0, 0; right < len(freeTime); right++ {
		sum += freeTime[right]
		if right-left+1 < k+1 {
			continue
		}
		ans = max(ans, sum)
		sum -= freeTime[left]
		left++
	}
	return ans
}

func minSwaps(nums []int) int {
	totalCnt := 0
	for _, num := range nums {
		if num == 1 {
			totalCnt++
		}
	}
	ans := totalCnt
	sum := 0
	for left, right := 0, 0; right < len(nums)+totalCnt; right++ {
		if nums[right%len(nums)] == 0 {
			sum++
		}
		if right-left+1 < totalCnt {
			continue
		}
		ans = min(ans, sum)
		if nums[left%len(nums)] == 0 {
			sum--
		}
		left++
	}
	return ans
}

func decrypt(code []int, k int) []int {
	codeLen := len(code)
	ans := make([]int, codeLen)
	if k == 0 {
		return ans
	}
	if k > 0 {
		sum := 0
		for left, right := 0, 0; left < codeLen; right++ {
			sum += code[right%codeLen]
			if right-left < k {
				continue
			}
			sum -= code[left]
			ans[left] = sum
			left++
		}
	} else {
		k = -k
		sum := 0
		for left, right := codeLen-k, codeLen-k; right < codeLen+codeLen+1; right++ {
			sum += code[right%codeLen]
			if right-left+1 < k {
				continue
			}
			ans[(right+1)%codeLen] = sum
			sum -= code[left%codeLen]
			left++
		}
	}
	return ans
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	records := make(map[byte]int, 26)
	ans := make(map[string]int)
	for left, right := 0, 0; right < len(s); right++ {
		records[s[right]]++
		if right-left+1 == minSize {
			if len(records) <= maxLetters {
				ans[s[left:right+1]]++
			}
			records[s[left]]--
			if records[s[left]] == 0 {
				delete(records, s[left])
			}
			left++
		}
	}
	maxNum := 0
	for _, v := range ans {
		maxNum = max(v, maxNum)
	}
	return maxNum
}

func minFlips(s string) int {
	n := len(s)
	min010 := n
	sum010 := 0
	for left, right := 0, 0; right < 2*n; right++ {
		if right%2 == 0 && s[right%n] == '1' {
			sum010++
		}
		if right%2 == 1 && s[right%n] == '0' {
			sum010++
		}
		if right-left+1 < n {
			continue
		}
		min010 = min(min010, sum010, n-sum010)
		if left%2 == 0 && s[left%n] == '1' {
			sum010--
		}
		if left%2 == 1 && s[left%n] == '0' {
			sum010--
		}
		left++
	}
	return min(min010, n-min010)
}

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)
	if n1 > n2 {
		return false
	}
	cnt1 := [26]int{}
	for _, ch := range s1 {
		cnt1[ch-'a']++
	}
	cnt2 := [26]int{}
	for left, right := 0, 0; right < n2; right++ {
		cnt2[s2[right]-'a']++
		if right-left+1 < n1 {
			continue
		}
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[left]-'a']--
		left++
	}
	return false
}

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	p := 1
	hash := 0
	n := len(s)
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + int(s[i]-'a'+1)) % modulo
		if i != n-k {
			p = p * power % modulo
		}
	}
	ans := ""
	if hash == hashValue {
		ans = s[n-k:]
	}
	for i := n - k - 1; i >= 0; i-- {
		hash = ((hash-int(s[i+k]-'a'+1)*p%modulo+modulo)*power + int(s[i]-'a'+1)) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return ans
}

func queryString(s string, n int) bool {
	numCnt := make(map[int]struct{})
	for i := 0; i < len(s); i++ {
		x := int(s[i] - '0')
		if x == 0 {
			continue
		}
		for j := i + 1; x <= n; j++ {
			numCnt[x] = struct{}{}
			if j == len(s) {
				break
			}
			x = x*2 + int(s[j]-'0')
		}
	}
	return n == len(numCnt)
}

func findAnagrams(s string, p string) []int {
	lenS := len(s)
	lenP := len(p)
	ans := []int{}
	if lenS < lenP {
		return ans
	}
	cntP := [26]int{}
	for i := 0; i < lenP; i++ {
		cntP[p[i]-'a']++
	}
	cntS := [26]int{}
	for left, right := 0, 0; right < lenS; right++ {
		cntS[s[right]-'a']++
		if right-left+1 < lenP {
			continue
		}
		if cntS == cntP {
			ans = append(ans, left)
		}
		cntS[s[left]-'a']--
		left++
	}
	return ans
}
