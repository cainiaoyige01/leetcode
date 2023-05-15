package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/13 10:37
 * @Description:最后一个单词的长度
 */
func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

// lengthOfLastWord 想分割字符串 注意split与Fields的区别 前者是按字符串进行分割 后者是按照空格进行分割
func lengthOfLastWord(str string) int {
	//先分割字符串先
	split := strings.Fields(str)
	if len(split) == 0 {
		return 0
	}
	s := split[len(split)-1]
	return len(s)
}
