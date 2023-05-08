package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/27 22:22
 * @Description:最长回文子串
 */
func main() {
	fmt.Println(longestPalindrome1("babad"))
}

// longestPalindrome1 使用动态规划来做！
func longestPalindrome1(str string) string {
	n := len(str)
	if n == 1 || n == 0 {
		return str
	}
	left, right := 0, n
	for left < right {
		if isPalindrome(str[left:right]) {
			return str[left:right]
		}
		left++
		right--
	}
	return str[left:right]
}

// longestPalindrome 这是暴力破解
func longestPalindrome(str string) string {
	if len(str) == 1 || len(str) == 0 {
		return str
	}
	max := ""
	n := len(str)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if len(str[i:j]) > len(max) && isPalindrome(str[i:j]) {
				max = str[i:j]
			}
		}

	}
	return max
}
func isPalindrome(str string) bool {
	n := len(str)
	left, right := 0, n-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
