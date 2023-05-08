package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/16 22:32
 * @Description:
 */
func main() {
	v := []int{1, 2, 3}
	fmt.Println(v)
	//在开始的时候就确定了切片的长度 也就是遍历的次数了
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}
