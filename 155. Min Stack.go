package main

import "fmt"

//TODO 逻辑没问题，但是通过不了越界的测试
type MinStack struct {
	cur  int
	min  []int
	data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	min, data := []int{}, []int{}
	return MinStack{
		cur:  -1,
		min:  min,
		data: data,
	}
}

func (this *MinStack) Push(x int) {
	this.cur++
	this.data = append(this.data, x)
	if this.cur > 0 {
		if x > this.min[this.cur-1] {
			this.min = append(this.min, this.min[this.cur-1])
		} else {
			this.min = append(this.min, x)
		}
	} else {
		this.min = append(this.min, x)
	}
	fmt.Println(this.min)
}

func (this *MinStack) Pop() {
	this.cur--
}

func (this *MinStack) Top() int {
	return this.data[this.cur]
}

func (this *MinStack) GetMin() int {
	return this.min[this.cur]
}

func main() {
	m := Constructor()
	m.Push(-2)
	m.Push(0)
	m.Push(-1)
}
