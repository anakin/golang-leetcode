package main

/**
如果我们能知道该数最高位的1所在的位置，就可以构造一个长度和该数据所占位置一样长的一个掩码mask，然后概述和mask进行异或即可。
*/

func findComplement(num int) int {
	mask := 1
	tmp := num
	for tmp > 0 {
		mask <<= 1
		tmp >>= 1
	}
	return num ^ (mask - 1)
}
