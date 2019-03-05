package main

import (
	"sort"
	"strconv"
)

//用内置排序接口
type intSlice []int

func newIntSlice(a []int) intSlice {
	b := intSlice{}
	for _, v := range a {
		b = append(b, v)
	}
	return b
}
func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Less(i, j int) bool {
	stri := strconv.Itoa(s[i])
	strj := strconv.Itoa(s[j])
	s1, _ := strconv.Atoi(stri + strj)
	s2, _ := strconv.Atoi(strj + stri)
	return s1 > s2
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	s := newIntSlice(nums)
	sort.Sort(s)
	res := ""
	for _, v := range s {
		res += strconv.Itoa(v)
		if res == "0" {
			res = ""
		}
	}
	if res == "" {
		return "0"
	}
	return res
}
