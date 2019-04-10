package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归把节点的数量保存在map中,同时记录最大的出现次数
*/

var mode int

func findMode(root *TreeNode) []int {
	mode = 0
	dic := make(map[int]int)
	ret := []int{}
	helper(root, dic)
	for k, v := range dic {
		if v == mode {
			ret = append(ret, k)
		}
	}
	return ret
}

func helper(root *TreeNode, dic map[int]int) {
	if root == nil {
		return
	}
	dic[root.Val]++
	if dic[root.Val] > mode {
		mode = dic[root.Val]
	}
	helper(root.Left, dic)
	helper(root.Right, dic)
}
