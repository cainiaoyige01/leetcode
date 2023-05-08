package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/20 0:40
 * @Description:
 */
func change(s ...int) {
	//s = append(s, 3)
	s[1] = 10
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	change(nums...)
	//f(nums)
	fmt.Println(nums)
}
func f(nums []int) {

}
