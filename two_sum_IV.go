package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	dataSet := make(map[int]struct{})
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			ele := q[0]
			if _, ok := dataSet[k-ele.Val]; ok {
				return true
			}
			q = q[1:]
			dataSet[ele.Val] = struct{}{}
			if ele.Right != nil {
				q = append(q, ele.Right)
			}
			if ele.Left != nil {
				q = append(q, ele.Left)
			}
		}
	}
	return false
}
