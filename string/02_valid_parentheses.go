package main

import "fmt"

/**
 * @Author: _niuzai
 * @Date:   2023/5/9 10:35
 * @Description:有效的括号
 */
func main() {
	fmt.Println(isValid("["))
	fmt.Println(isValid("([)]"))
}

// isValid
func isValid(str string) bool {
	//rune 类似与int32 会自动按照字符独立的单位处理
	var runes []rune
	for _, v := range str {
		switch v {
		case '{', '(', '[':
			runes = append(runes, v)
		case '}', ')', ']':
			if len(runes) > 0 && opposite(v) == runes[len(runes)-1] {
				//这是秒点之一 缩小组数 就是弹掉的那一部门
				runes = runes[0 : len(runes)-1]
			} else {
				return false
			}
		}
	}
	if len(runes) > 0 {
		return false
	}
	return true
}

// opposite 目的是为了反弹出去 也就是类似与弹栈
func opposite(v int32) rune {
	switch v {
	case '}':
		return '{'
	case ')':
		return '('
	case ']':
		return '['
	}
	return '\n'
}
