package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i := 0; cnt < n && i < len(flowerbed); i++ {
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}
		if i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			continue
		}
		if flowerbed[i] != 0 {
			continue
		}
		flowerbed[i] = 1
		cnt++
		if i < len(flowerbed)-1 {
			flowerbed[i+1] = -1
		}
	}
	return cnt >= n
}
