package main

import (
	"fmt"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/15 16:51
 * @Description:检测大写字母
 */
func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Leetcode"))
	fmt.Println(detectCapitalUse("mL"))
}

/**
官方做法：  若第 1 个字母为小写，则需额外判断第 2 个字母是否为小写
		   若第 1 个字母为小写，则需额外判断第 2 个字母是否为小写
*/
// detectCapitalUse 使用额外数组来做
func detectCapitalUse(s string) bool {
	if len(s) == 0 {
		return true
	}
	count := 0
	flag := false
	if int32(s[0]) >= 'A' && int32(s[0]) <= 'Z' {
		flag = true
	}
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			count++
		}
	}
	if flag && count == 1 || count == len(s) {
		return true
	}
	return false

}
