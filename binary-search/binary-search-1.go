package binary_search

func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}

func searchInsert(nums []int, target int) int {
	return lowerBound(nums, target)
}

func search(nums []int, target int) int {
	idx := lowerBound(nums, target)
	if idx == len(nums) || nums[idx] != target {
		return -1
	}
	return idx
}

func lowerBoundByte(letters []byte, target byte) int {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := lowerBoundByte(letters, target+1)
	return letters[idx%len(letters)]
}

func maximumCount(nums []int) int {
	num1 := lowerBound(nums, 0)
	num2 := len(nums) - lowerBound(nums, 1)
	return max(num1, num2)
}
