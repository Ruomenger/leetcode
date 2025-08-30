package btree

import (
	"container/list"
)

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

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	stack := list.New()
	node := root
	var prev *TreeNode
	for node != nil || stack.Len() != 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}
		ele := stack.Back()
		stack.Remove(ele)
		node = ele.Value.(*TreeNode)

		if node.Right == nil || node.Right == prev {
			ans = append(ans, node.Val)
			prev = node
			node = nil
		} else {
			stack.PushBack(node)
			node = node.Right
		}
	}

	return ans
}

func getLeaf(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
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
		if node.Left == nil && node.Right == nil {
			ans = append(ans, node.Val)
		}
		node = node.Right
	}
	return ans
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := getLeaf(root1)
	leaf2 := getLeaf(root2)
	if len(leaf1) != len(leaf2) {
		return false
	}
	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}
