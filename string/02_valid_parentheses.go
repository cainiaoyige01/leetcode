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

func isValid(s string) bool {
	//定义一个切片来存储右边的括号 rune 可以确定每一个元素
	var str []rune
	for _, v := range s {
		switch v {
		case '[', '{', '(':
			//存储到str中去
			str = append(str, v)
		case ']', '}', ')':
			//首先要判断切片中是否有元素 然后是该元素是否与反向的相等
			if len(str) > 0 && apposite(v) == str[len(str)-1] {
				//进行类似与栈进行弹栈
				str = str[:len(str)-1]
			} else {
				return false
			}
		}
	}
	//判断str中是否还有元素
	if len(str) > 0 {
		return false
	}

	return true
}
func apposite(v int32) rune {
	switch v {
	case ']':
		return '['
	case '}':
		return '{'
	case ')':
		return '('
	}
	return '\n'
}
