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

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := list.New()
	node := root
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		ele := stack.Back()
		stack.Remove(ele)
		node = ele.Value.(*TreeNode)
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans
}
