package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	nn := []int{}
	if len(nums) <= 1 {
		nn = nums
	} else {
		nn = append(nn, nums[0])
		for i := 1; i < len(nums); i++ {
			nn = append(nn, nums[i]+nn[i-1])
		}
	}
	return NumArray{
		nums: nn,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.nums[j]
	} else {
		return this.nums[j] - this.nums[i-1]
	}
}

//
//func main() {
//	r := []int{-2, 0, 3, -5, 2, -1}
//	a := Constructor(r)
//	b := a.SumRange(0, 2)
//	fmt.Println(b)
//}
