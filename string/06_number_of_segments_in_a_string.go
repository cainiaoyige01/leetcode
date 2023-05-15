package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/15 16:43
 * @Description:字符串中的单词数
 */
func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}
func countSegments(s string) int {
	return len(strings.Fields(s))
}
