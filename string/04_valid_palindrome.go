package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/14 11:23
 * @Description:验证回文串
 */
func main() {
	fmt.Println(isPalindrome1("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome1("abba"))

}
func reverseString2(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

// isPalindrome1 更加便捷的方法是把是字母或数字的 全新组成一组字符串 ！然后再来求会文串！不过这样会消耗空间时间
func isPalindrome1(s string) bool {
	//首相第一步是把里面的字母全变成小写
	lower := strings.ToLower(s)
	fmt.Println(lower)
	//使用双指针进行便利
	left := 0
	right := len(s) - 1
	//循环遍历
	for left < right {
		//在这里可能会遇到多个非字母或数字的字符 需要多次进行跳过了
		for left < right && !(lower[left] >= 'a' && lower[left] <= 'z' || lower[left] >= '0' && lower[left] <= '9') {
			left++
		}
		for left < right && !(lower[right] >= 'a' && lower[right] <= 'z' || lower[right] >= '0' && lower[right] <= '9') {
			right--
		}
		if lower[left] == lower[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
