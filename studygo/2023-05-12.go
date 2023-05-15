package main

import "fmt"

/**
 * @Author: _niuzai
 * @Date:   2023/5/12 21:31
 * @Description:
 */
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//这里通过不了编译 需要进行语法糖进行拆分
	//s1 = append(s1, s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
