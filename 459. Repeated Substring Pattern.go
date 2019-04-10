package main

/**
有一个神奇的解法：
1. 用两个s 首尾相连得到一个新的字符串ss ;
2. 去掉ss 的首尾两个字符；
3. 如果在剩下来的字符串中能找到s 那么返回True，否则False
if len(s) == 0 {
        return false
    }
    ss := (s + s)[1:(len(s)*2)-1]
    return strings.Contains(ss, s)
这里用的解法：
重复子串长度最长为len/2，直接每次选择一个可以被整除的较小的数，截取开头的那几个字符，然后重复到原长度验证就好。
*/
func repeatedSubstringPattern(s string) bool {
	str := []byte(s)
	if len(str) <= 1 {
		return false
	}
	max := len(str) / 2
	for i := 1; i <= max; i++ {
		j := i
		idx := 0
		if len(str)%i > 0 {
			continue
		}
		for ; j < len(str); j++ {
			if str[j] != str[idx] {
				break
			}
			idx++
			if idx >= i {
				idx = 0
			}
		}

		if j == len(str) {
			return true
		}
	}
	return false
}
