package main

import "sort"

/**
将区间排序，因为我要们除掉的区间尽可能小，所以我尽量保存小区间，去除大区间，所以我们以右区间排序，从小到大，并且右区间相等时，按左区间从小到大排序，然后我们就维护左区间边界，当左区间边界小于下一个的区间的右边界的时候，就将左区间边界扩张为一个区间的左边界，当左区间边界大于下一个的区间的右边界的时候，我们要比较左区间边界与下一个区间的左边界大小，取最小的，这样我们可以容纳更多的区间而不重叠。
*/

type Interval struct {
	Start int
	End   int
}

type inSlice []Interval

func newInSlice(a []Interval) inSlice {
	b := inSlice{}
	for _, v := range a {
		b = append(b, v)
	}
	return b
}

func (this inSlice) Len() int {
	return len(this)
}
func (this inSlice) Less(i, j int) bool {
	if this[i].Start == this[j].Start {
		return this[i].End < this[j].End
	} else {
		return this[i].Start < this[j].Start
	}
}
func (this inSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) < 2 {
		return 0
	}
	in := newInSlice(intervals)
	sort.Sort(in)
	left := in[0].End
	cnt := 1
	for i := 1; i < len(in); i++ {
		if left <= in[i].Start {
			cnt++
			left = in[i].End
		} else {
			if left > in[i].End {
				left = in[i].End
			}
		}
	}
	return len(in) - cnt
}
