package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i, j := 0, len(nums)-1; i < j; i++ {
		if nums[i] != nums[j] {
			return false
		}
		j--
	}
	return true
}
