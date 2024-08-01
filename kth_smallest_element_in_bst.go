package main

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	current := root
	for {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k = k - 1
		if k == 0 {
			return current.Val
		}
		current = current.Right
	}
}
