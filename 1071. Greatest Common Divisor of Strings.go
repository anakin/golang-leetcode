/**
For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T concatenated with itself 1 or more times)

Return the largest string X such that X divides str1 and X divides str2.



Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

*/
package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/strings"

func gcdOfStrings(str1 string, str2 string) string {

	if len(str1) == 0 || str1 == str2 {
		return str1
	}
	if len(str2) == 0 {
		return str2
	}
	if len(str1) > len(str2) {
		x := strings.Index(str1, str2)
		if x > -1 {
			return gcdOfStrings(str2, str1[x+len(str2):])
		}
	} else {
		y := strings.Index(str2, str1)
		if y > -1 {
			return gcdOfStrings(str1, str2[y+len(str1):])
		}
	}
	return ""
}

/**
need to know the function:Strings.Index
*/
