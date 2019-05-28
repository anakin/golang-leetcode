package main

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string s, int numRows);
Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

*/

/**
关键是要算出循环的长度：
为 row + row -2
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	item_len := 2*numRows - 2 //循环的长度
	res := make([][]string, numRows, numRows)

	for index, v := range s {

		mod := index % item_len

		if mod < numRows {
			res[mod] = append(res[mod], string(v))
		} else {
			i := numRows - (mod - numRows) - 2
			res[i] = append(res[i], string(v))
		}
	}

	var str string

	for _, arr := range res {
		for _, v := range arr {
			str += v
		}
	}

	return str
}
