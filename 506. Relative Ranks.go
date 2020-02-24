/*
Given scores of N athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
For the left two athletes, you just need to output their relative ranks according to their scores.

*/
package main

func findRelativeRanks(nums []int) []string {
	sorted:=make([]int,len(nums))
	copy(sorted,nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	places:=map[int]int{}
	for i:=0;i<len(sorted);i++{
		places[sorted[i]] = i+1
	}
	var res []string
	for _,num:=range nums{
		p:=places[num]
		if p == 1{
			res = append(res,"Gold Medal")
		}else if p == 2{
			res = append(res,"Silver Medal")
		}else if p == 3{
			res = append(res,"Bronze Medal")
		}else{
			res = append(res,strconv.Itoa(p))
		}

	}
	return res
}

/*
先复制出一份，按照从达到小排好序，然后从原来的slice中搜索对应的位置进行替换
 */
