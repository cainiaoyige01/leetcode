package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/9 0:33
 * @Description:最长公共前缀
 */
func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "fly", "flight"})) // fl
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))   // ""
}

// longestCommonPrefix  维护一个前缀库 不断往下遍历 判断前缀然乎开始裁剪前缀库的长度
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	s1 := strs[0]
	prefix := make([]string, len(s1))
	for i := range s1 {
		prefix[i] = s1[:i+1]
	}
	for _, str := range strs {
		for i2 := range prefix {
			//这是秒点 就是利用了 str[0:i2]==prefix[i2]的原理
			if !strings.HasPrefix(str, prefix[i2]) {
				prefix = prefix[:i2]
				break
			}
		}
	}

	if len(prefix) > 0 {
		return prefix[len(prefix)-1]
	}
	return ""
}
