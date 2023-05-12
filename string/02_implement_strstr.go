package main

import (
	"fmt"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/12 16:10
 * @Description:实现strStr()
 */
func main() {
	fmt.Println(strStr("a", ""))
	fmt.Println(strStr("hello", "ll"))
}

// strStr
func strStr(str1 string, str2 string) int {
	l1 := len(str1)
	l2 := len(str2)
	if l2 == 0 {
		return 0
	}
	if l2 > l1 {
		return -1
	}
	var str3 []int
	s := int32(str2[0])
	for i, v := range str1 {
		if s == v {
			str3 = append(str3, i)
		}
	}
	for _, v := range str3 {
		if v+l2 > l1 {
			break
		}
		if string(str1[v:v+l2]) == str2 {
			return v
		}
	}
	return -1
}
