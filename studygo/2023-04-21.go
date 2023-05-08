package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/21 23:40
 * @Description:
 */
func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int
	//和20的题一样切片内部是一个struct包含一个数组指针 只要不发生扩容 里面的元素都会发生修改
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}
