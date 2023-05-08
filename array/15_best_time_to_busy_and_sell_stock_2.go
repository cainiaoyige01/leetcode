package main

import (
	"fmt"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/18 10:12
 * @Description:买卖股票的最佳时机||
 */
func main() {
	fmt.Println(maxProfit1([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit1(nums []int) int {
	n := len(nums)
	i := 0
	sum := 0
	for i < n-1 {
		buy := i
		for i < n-1 && nums[i] < nums[i+1] {
			i++
		}
		//在最高点进行抛售
		sum += nums[i] - nums[buy]
		i++
	}
	return sum
}
