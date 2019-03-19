package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
用队列保存节点,一层一层的处理
从队列里取出节点的同时，判断左右子树是否为空，不为空则入队列
*/

func averageOfLevels1(root *TreeNode) []float64 {
	ret := []float64{}
	if root == nil {
		return ret
	}
	cnt, sum := 0, 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		cnt = len(q)
		sum = 0
		for i := 0; i < cnt; i++ {
			tmp := q[0]
			q = q[1:]
			if tmp.Left != nil {
				q = append(q, tmp.Left)
			}
			if tmp.Right != nil {
				q = append(q, tmp.Right)
			}
			sum += tmp.Val
		}
		ret = append(ret, float64(sum)/float64(cnt))
	}
	return ret
}
