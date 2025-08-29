package btree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		ele := stack.Back()
		stack.Remove(ele)
		node := ele.Value.(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ans
}
