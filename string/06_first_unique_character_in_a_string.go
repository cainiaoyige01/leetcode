package main

import "fmt"

/**
 * @Author: _niuzai
 * @Date:   2023/5/15 16:28
 * @Description:字符串中的第一个唯一字符
 */
func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

// firstUniqChar  利用map rune来做
func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}
	return -1
}
