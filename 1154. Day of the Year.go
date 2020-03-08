/**
Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.



Example 1:

Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
Example 2:

Input: date = "2019-02-10"
Output: 41

*/
package main

import "strconv"

func dayOfYear(date string) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		months[1] = 29
	}
	return sum(months, month-1) + day
}

func sum(arr []int, index int) int {
	var res int
	for i := 0; i < index; i++ {
		res += arr[i]
	}
	return res
}
