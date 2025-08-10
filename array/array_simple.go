package array

func twoSum1(nums []int, target int) []int {
	ans := []int{}
OUT:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
				break OUT
			}
		}
	}
	return ans
}

func twoSum2(nums []int, target int) []int {
	ans := []int{}
	posMap := make(map[int][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		posMap[nums[i]] = append(posMap[nums[i]], i)
	}
OUT:
	for i := 0; i < len(nums); i++ {
		if pos, ok := posMap[target-nums[i]]; ok {
			for j := range pos {
				if pos[j] == i {
					continue
				}
				ans = append(ans, i, pos[j])
				break OUT
			}
		}
	}
	return ans
}

func twoSum3(nums []int, target int) []int {
	num2IdxMap := make(map[int]int, len(nums))
	ans := make([]int, 0, 2)
	for i, num := range nums {
		if idx, ok := num2IdxMap[target-num]; ok {
			ans = append(ans, idx, i)
			break
		}
		num2IdxMap[num] = i
	}
	return ans
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return float64(findKthNum(nums1, nums2, totalLen/2+1))
	}
	n1 := findKthNum(nums1, nums2, totalLen/2)
	n2 := findKthNum(nums1, nums2, totalLen/2+1)
	return float64(n1+n2) / 2
}

func findKthNum(nums1, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		return findKthNum(nums2, nums1, k)
	}

	m := len(nums1)
	if m == 0 {
		return nums2[k-1]
	}

	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	i := min(k/2, m)
	j := k - i

	if nums1[i-1] < nums2[j-1] {
		return findKthNum(nums1[i:], nums2, k-i)
	} else {
		return findKthNum(nums1, nums2[j:], k-j)
	}
}

func maxArea(height []int) int {
	ans := 0
	for left, right := 0, len(height)-1; left < right; {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
