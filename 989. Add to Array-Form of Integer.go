/*
For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.



Example 1:

Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234

 */
package main

func addToArrayForm(A []int, K int) []int {

	i := len(A) - 1
	res := []int{}

	for i >= 0 || K > 0 {
		if i >= 0 {
			K += A[i]
		}
		res = append(res, K%10)
		K /= 10
		i--
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
/*
note:
最开始想到的是先把A变成一个整数，然后和K相加，然后把结果再变成数组，但是跑了一下用例，发现这样会溢出

 */