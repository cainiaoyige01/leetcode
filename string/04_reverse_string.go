package main

import "fmt"

/**
 * @Author: _niuzai
 * @Date:   2023/5/14 11:44
 * @Description:反转字符串
 */
func main() {
	fmt.Println(reverseString("hello"))
}

func reverseString(s string) string {
	lenS := len(s)
	//使用rune
	runes := make([]rune, lenS)
	for i := 0; i < lenS; i++ {
		runes[i] = rune(s[lenS-1-i])
	}
	return string(runes)
}
