package main

/**

设置一个队列
第一步现将所有入度为0的节点保存在队列中
然后，出队列，并将该节点的后续节点的入度减一
如果后续节点的入度变成0，入队列
最后将结果逆序
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	m := make([][]int, numCourses)
	in := make([]int, numCourses)
	q := []int{}
	order := []int{}
	for _, p := range prerequisites {
		m[p[0]] = append(m[p[0]], p[1]) //记录所有的对应关系
		in[p[1]]++                      //记录入度
	}
	//把所有入度为0的节点，入队列
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	//循环处理，一边出队列，一边入队列
	for len(q) != 0 {
		cur := q[0]
		q = q[1:] //出队列
		order = append(order, cur)
		//处理该节点的后续节点
		for i := 0; i < len(m[cur]); i++ {
			dep := m[cur][i]
			in[dep]-- //后续节点-1
			if in[dep] == 0 {
				//如果入度变成了0，入队列
				q = append(q, dep)
			}
		}
	}
	//如果还有节点没有处理，则不能完成课程
	if len(order) != numCourses {
		return []int{}
	}
	//逆序
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}
