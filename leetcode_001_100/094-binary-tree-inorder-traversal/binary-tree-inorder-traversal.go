package problem0094

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return []int{root.Val}
//	}
//
//	res := inorderTraversal(root.Left)
//	res = append(res, root.Val)
//	res = append(res, inorderTraversal(root.Right)...)
//
//	return res
//}

// stack
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	p := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		// 左节点全部压栈，然后遍历
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		fmt.Println("curr",curr)
		p = curr.Right
	}
	fmt.Println("res", res)
	return res
}

// stack  pre
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	p := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) > 0 || p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		// 左节点全部压栈，然后遍历
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		fmt.Println("curr",curr)
		p = curr.Right
	}
	fmt.Println("res", res)
	return res
}