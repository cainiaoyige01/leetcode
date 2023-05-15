package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/14 12:36
 * @Description:反转字符串中的元音字母
 */
func main() {
	fmt.Println(reverseVowels("ai"))
}
func reverseVowels(s string) string {
	runes := []rune(s)
	left := 0
	right := len(runes) - 1
	//其实可以使用strings.Contains("aeiouAEIOU",string(s[i]))来判断是否包含
	for left < right {
		temp := strings.ToLower(string(s[left]))
		for left < right && !(temp == "a" || temp == "i" || temp == "e" || temp == "o" || temp == "u") {
			left++
			temp = strings.ToLower(string(s[left]))
		}
		temp = strings.ToLower(string(s[right]))
		for left < right && !(temp == "a" || temp == "i" || temp == "e" || temp == "o" || temp == "u") {
			right--
			temp = strings.ToLower(string(s[right]))
		}
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}
	return string(runes)
}
