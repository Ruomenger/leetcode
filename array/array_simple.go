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
