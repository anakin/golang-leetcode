package main

/**
从叶子节点开始不停的向中间收缩。
首先遍历所有的节点，保存对应的入度关系（由于没有方向，所以需要保存两个方向的关系）
然后找到所有的叶子节点（入度等于1的节点）
然后先将叶子节点入队列
剪掉这些叶子节点的后续节点，同时将其中入度为1的节点入队列
当叶子节点少于三个的时候,就是最后的结果

*/
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	m := make(map[int][]int)
	//建立对应关系
	for _, v := range edges {
		if _, ok := m[v[0]]; ok {
			m[v[0]] = append(m[v[0]], v[1])
		} else {
			m[v[0]] = []int{v[1]}
		}
		if _, ok := m[v[1]]; ok {
			m[v[1]] = append(m[v[1]], v[0])
		} else {
			m[v[1]] = []int{v[0]}
		}
	}
	leaves := []int{}
	//找到所有的叶子节点
	for k, v := range m {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}
	for n > 2 {
		n -= len(leaves)
		one := []int{}
		for _, v := range leaves {
			next := m[v]
			for _, val := range next {
				m[v] = del(m[v], val)
				m[val] = del(m[val], v)
				if len(m[val]) == 1 {
					one = append(one, val)
				}
			}
		}
		leaves = one
	}
	return leaves
}

func del(s []int, v int) []int {
	for key, val := range s {
		if val == v {
			s = append(s[:key], s[key+1:]...)
		}
	}
	return s
}
