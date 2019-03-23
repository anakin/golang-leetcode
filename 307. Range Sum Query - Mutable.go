package main

type NumArray struct {
	n []int
}

func Constructor(nums []int) NumArray {
	n := []int{}
	if len(nums) <= 1 {
		n = nums
	} else {
		n = append(n, nums[0])
		for i := 1; i < len(nums); i++ {
			n = append(n, nums[i]+n[i-1])
		}
	}
	return NumArray{
		n: n,
	}
}

func (this *NumArray) Update(i int, val int) {
	l := len(this.n)
	if i < 0 || i > l-1 {
		return
	}
	var diff int
	if i == 0 {
		diff = val - this.n[0]
	} else {
		diff = val - (this.n[i] - this.n[i-1])
	}
	for j := i; j < l; j++ {
		this.n[j] += diff
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.n[j]
	} else {
		return this.n[j] - this.n[i-1]
	}
}
