package main

import (
	"sort"
)

/**
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	l := len(intervals)
	var (
		starts []int
		ends   []int
		res    []Interval
	)
	for _, v := range intervals {
		starts = append(starts, v.Start)
		ends = append(ends, v.End)
	}
	sort.Ints(starts)
	sort.Ints(ends)
	for i, j := 0, 0; i < l; i++ {
		if i == l-1 || starts[i+1] > ends[i] { //中间出现了不连续，记录前一个连续值
			res = append(res, Interval{Start: starts[j], End: ends[i]})
			j = i + 1
		}
	}
	return res
}
