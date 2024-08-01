package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
   if p==nil && q==nil {
   	  return true
   }
	if p==nil || q==nil {
		return false
	}
	if p.Val!=q.Val{
		return false
	}
	return isSameTree(p.Right,q.Right) && isSameTree(p.Left,q.Left)
}



func LevelOrderTraversal(root *TreeNode) [][]int{
	var result = make([][]int,0)
	if root==nil{
		return [][]int{}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue)>0{
		tempResult := make([]int,0)
		size := len(queue)
		for i:= 1;i<=size;i++{
			ele := queue[0]
			queue = queue[1:]
			tempResult = append(tempResult, ele.Val)
			if ele.Left!=nil{
				queue = append(queue, ele.Left)
			}
			if ele.Right!=nil{
				queue = append(queue, ele.Right)
			}
		}
		result = append(result,tempResult)
	}
	return result
}

func ZigZaglevelOrderTraversal(root *TreeNode) [][]int{
	var result = make([][]int,0)
	if root==nil{
		return [][]int{}
	}
	var queue []*TreeNode
	queue = append(queue, root)
	toggleFlag := false
	for len(queue)>0{
		   tempResult := make([]int,0)
		   size := len(queue)
		   for i:= 1;i<=size;i++{
			    ele := queue[0]
			    queue = queue[1:]
			    if toggleFlag{
					tempResult = append([]int{ele.Val},tempResult...)
				}else{
					tempResult = append(tempResult, ele.Val)
				}
			   if ele.Left!=nil{
				   queue = append(queue, ele.Left)
			   }
			   if ele.Right!=nil{
				   queue = append(queue, ele.Right)
			   }
		   }
		   result = append(result,tempResult)
		   toggleFlag = !toggleFlag
	}
	fmt.Println(result)
	return result
}


func main() {
	// BST Representation
	//         100
	//        /   \
	//      50    150
	//     / \    /  \
	//   20  80  110 200
	//  /  \
	// 10  30

	root := TreeNode{100, nil, nil}
	root.Left = &TreeNode{50, nil, nil}
	root.Right = &TreeNode{150, nil, nil}
	root.Left.Left = &TreeNode{20, nil, nil}
	root.Left.Left.Left = &TreeNode{10, nil, nil}
	root.Left.Left.Right = &TreeNode{30, nil, nil}
	root.Left.Right = &TreeNode{80, nil, nil}
	root.Right.Left= &TreeNode{110, nil, nil}
	root.Right.Right = &TreeNode{200, nil, nil}
    ZigZaglevelOrderTraversal(&root)
	LevelOrderTraversal(&root)

}