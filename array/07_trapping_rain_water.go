package main

import (
	"fmt"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/11 14:33
 * @Description:接雨水
 */
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}

// trap 利用每一格的高度差来做的！是一步一步来的
func trap(nums []int) int {
	area := 0
	//定义左右指针
	left, right := 0, len(nums)-1
	//定义起点高度 假设都为0
	lTop, rTop := 0, 0
	for left <= right {
		//在这里取高度 lTop以及rTop均是取最大的高度
		lTop = max(nums[left], lTop)
		rTop = max(nums[right], rTop)
		//那侧偏高 往那侧走
		if lTop < rTop {
			area += lTop - nums[left]
			left++
		} else {
			area += rTop - nums[right]
			right--
		}
	}
	return area

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
