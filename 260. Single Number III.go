package main

/**
如果能将A，B分开到二个数组中，那显然符合“异或”解法的关键点了。因此这个题目的关键点就是将A，B分开到二个数组中。由于A，B肯定是不相等的，因此在二进制上必定有一位是不同的。根据这一位是0还是1可以将A，B分开到A组和B组。而这个数组中其它数字要么就属于A组，要么就属于B组。再对A组和B组分别执行“异或”解法就可以得到A，B了。而要判断A，B在哪一位上不相同，只要根据A异或B的结果就可以知道了，这个结果在二进制上为1的位就说明A，B在这一位上是不相同的。
*/
func singleNumber(nums []int) []int {
	res := make([]int, 2)
	xor := nums[0]
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}
	bit := xor &^ (xor - 1)
	n1, n2 := 0, 0
	for _, v := range nums {
		if v&bit > 0 {
			n1 ^= v
		} else {
			n2 ^= v
		}
	}
	res[0] = n1
	res[1] = n2
	return res
}
