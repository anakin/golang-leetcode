package main

/**
二维slice保存数据
*/
type MyCalendar struct {
	cal [][]int
}

func Constructor() MyCalendar {
	s := [][]int{}
	return MyCalendar{
		cal: s,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, v := range this.cal {
		if v[0] < end && start < v[1] {
			return false
		}
	}
	s := []int{start, end}
	this.cal = append(this.cal, s)
	return true
}
