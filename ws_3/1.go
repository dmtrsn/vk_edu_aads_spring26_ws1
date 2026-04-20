package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	q := make([]*TreeNode, 0, 1)
	q = append(q, root)

	for len(q) > 0 {
		n := len(q)
		currLevel := make([]int, 0, n)

		for j := 0; j < n; j++ {
			cur := q[0]
			q = q[1:]

			currLevel = append(currLevel, cur.Val)

			if left := cur.Left; left != nil {
				q = append(q, left)
			}
			if right := cur.Right; right != nil {
				q = append(q, right)
			}
		}

		res = append(res, currLevel)
	}

	return res
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	result := levelOrder(root)

	fmt.Println(result)
}
