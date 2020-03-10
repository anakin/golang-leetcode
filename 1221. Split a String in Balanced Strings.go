/**
Balanced strings are those who have equal quantity of 'L' and 'R' characters.

Given a balanced string s split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.



Example 1:

Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.
Example 2:

Input: s = "RLLLLRRRLR"
Output: 3
Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

*/
package main

func balancedStringSplit(s string) int {
	var sum, res int
	for _, c := range s {
		if c == 'R' {
			sum += 1
		} else {
			sum -= 1
		}
		if sum == 0 {
			res += 1
		}
	}
	return res
}
