package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (arr [][]int) {
	dao := func(arr []int) []int {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		return arr
	}
	if root == nil {
		return arr
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		a := []int{}
		for _, v := range q {
			node := v
			a = append(a, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}

		}
		q = p
		// 表示是奇数, 实现倒序
		if len(arr)%2 != 0 {
			a = dao(a)
		}
		arr = append(arr, a)
	}
	return arr
}
func main() {
	a := 10
	if a < 0 {
		fmt.Println(a)
	} else {

	}
}
