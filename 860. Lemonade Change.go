package main

func lemonadeChange(bills []int) bool {
	m := make([]int, 2)
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			m[0]++
			break
		case 10:
			if m[0] <= 0 {
				return false
			}
			m[0]--
			m[1]++
			break
		case 20:
			flag := false
			if m[0] >= 1 && m[1] >= 1 {
				flag = true
				m[0]--
				m[1]--
			} else if m[0] >= 3 {
				flag = true
				m[0] -= 3
			}
			if !flag {
				return false
			}
			break
		}

	}
	return true
}
