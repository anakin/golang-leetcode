package main

/**
分别记录上下左右的数量，
如果上和下相等，并且左和右相等，则返回true
*/
func judgeCircle(moves string) bool {
	mm := []byte(moves)
	if len(moves) == 0 {
		return true
	}
	su, sd, sl, sr := 0, 0, 0, 0
	for _, v := range mm {
		if v == 'U' {
			su++
		}
		if v == 'D' {
			sd++
		}
		if v == 'L' {
			sl++
		}
		if v == 'R' {
			sr++
		}
	}
	return su == sd && sl == sr
}
