package main

/**
首先由于对称性，target是正是负影响不大。
当sum-target为偶数，1+...-（sum-target）/2+...+k=target，那么答案依然是k。#2
当sum-target为奇数，那么sum-target+1是一个偶数
类似#2的证明，1+...-(sum-target+1)/2+...k=target-1
此时再考虑k的奇偶性。
如果k是偶数并且k>sum-target+1
那么1+...-(sum-target+1)/2+....-(k/2)...+k+(k+1)=target
由#2相似可证，相当于在1+2....+k+(k+1)减去了sum-target+1和k。
等价于sum+（k+1）-sum+target-1-k====>target也就是答案是k+1.#3
如果k=sum-target+1，由#3可知依然是k+1.#4
如果k是奇数：
1+2+...-(sum-target+1)/2.....+k-(k+1)+(k+2)=sum-(sum-target+1)+1=target,
因此答案是k+2.#5
*/
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for target > 0 {
		k++
		target -= k
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}
