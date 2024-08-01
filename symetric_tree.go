package main

func isSymmetric(root *TreeNode) bool {
	p := root.Left
	q := root.Right

	if p==nil && q==nil {
		return true
	}
	if p==nil || q==nil {
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return isSameTree(p.Right,q.Left) && isSameTree(p.Left,q.Right)
}
