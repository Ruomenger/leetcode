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
