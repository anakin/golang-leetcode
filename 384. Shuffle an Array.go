/*
Shuffle a set of numbers without duplicates.

Example:

// Init an array with set 1, 2, and 3.
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3] must equally likely to be returned.
solution.shuffle();

// Resets the array back to its original configuration [1,2,3].
solution.reset();

// Returns the random shuffling of array [1,2,3].
solution.shuffle();

 */
package main
type Solution struct {
arr []int
}


func Constructor(nums []int) Solution {
return Solution{arr:nums}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
return this.arr
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
tmp:=make([]int,len(this.arr))
copy(tmp,this.arr)
for i:=0;i<len(tmp);i++{
	r:=rand.Intn(len(tmp))
	tmp[i],tmp[r] = tmp[r],tmp[i]
}
return tmp
}