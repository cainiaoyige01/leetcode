package main

import "fmt"

/**
 * @Author: _niuzai
 * @Date:   2023/5/15 12:53
 * @Description: 赎金信
 */
func main() {
	fmt.Println(canConstruct2("aab", "aba"))
}

func canConstruct(s1 string, s2 string) bool {
	lenS1 := len(s1)
	lenS2 := len(s2)
	if lenS2 < lenS1 {
		return false
	} else if s1 == s2 {
		return true
	}
	for i := 0; i < lenS2-lenS1; i++ {
		if s2[i:lenS1] == s1 {
			return true
		}
	}
	return false
}

// canConstruct2 使用哈希map来做
func canConstruct2(s1 string, s2 string) bool {
	m1 := make(map[int32]int, len(s1))
	for _, v := range s1 {
		m1[v] = m1[v] + 1
	}
	for _, v := range s2 {
		i, ok := m1[v]
		if ok && i != 0 {
			m1[v] = m1[v] - 1
		}
	}
	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}
	return true
}
