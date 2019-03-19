package main

//当糖果的种类数目大于妹妹应分得的糖果数目时，妹妹所能分得的最大糖果种类数为妹妹应分得的糖果数，反之，则为糖果的种类数。
func distributeCandies(candies []int) int {
	m := make(map[int]int)
	for _, v := range candies {
		if _, ok := m[v]; !ok {
			m[v] = 1
		}
	}
	s := len(candies) / 2
	t := len(m)
	if t > s {
		return s
	} else {
		return t
	}

}
