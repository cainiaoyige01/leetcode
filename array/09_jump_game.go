package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/13 9:50
 * @Description:跳跃游戏
 */
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}, true)) // true
	fmt.Println(jump([]int{3, 2, 1, 0, 4}, true)) // false
	fmt.Println(jump([]int{1, 2, 3}, true))       // true
	fmt.Println(jump([]int{0, 1}, true))          // false
	fmt.Println(jump([]int{0, 2, 3}, true))       // false
	fmt.Println(jump([]int{2, 0, 0}, true))       // true
	fmt.Println(jump([]int{2, 0, 1, 0, 1}, true)) // false
}

// jump  关键点是遍历一次 寻找一个元素能到达的最远的点并且元素下标不能超过该点 也就是贪心算法
func jump(nums []int, b bool) bool {
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		//i跑到max后面去了 说明没有最大的值可以到达最后一位元素
		if i > max {
			return false
		}
		temp := nums[i] + i
		if temp > max {
			max = temp
		}
	}
	return true
}
