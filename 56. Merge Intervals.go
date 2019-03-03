package main

import (
	"sort"
)

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
