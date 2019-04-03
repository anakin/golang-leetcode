package main

/**
DAG，顾名思义，如果一个有向图中从任意点出发都不能回到该点的话，这张图就是一个有向无环图
判断一个有向图是否有环，有两个算法：

拓扑排序
即找出该图的一个线性序列，使得需要事先完成的事件总排在之后才能完成的事件之前。如果能找到这样一个线性序列，则该图是一个有向无环图
DFS
遍历图中的所有点，从i点出发开始DFS，如果在DFS的不断深入过程中又回到了该点，则说明图中存在回路。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int][]int)
	res := make([]int, numCourses)
	for _, p := range prerequisites {
		v, ok := m[p[0]]
		if !ok {
			m[p[0]] = []int{p[1]}
		} else {
			m[p[0]] = append(v, p[1])
		}
	}
	for i := 0; i < numCourses; i++ {
		if helper(i, res, m) == false {
			return false
		}
	}
	return true
}
func helper(idx int, res []int, m map[int][]int) bool {
	if res[idx] == 1 {
		return true
	}
	if res[idx] == 2 {
		return false
	}
	v, ok := m[idx]
	if !ok {
		return true
	}
	res[idx] = 2
	for i := 0; i < len(v); i++ {
		if helper(v[i], res, m) == false {
			return false
		}
	}
	res[idx] = 1
	return true
}
