package main

import (
	"fmt"
	"math"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/17 22:44
 * @Description:买卖股票的最佳时机
 */
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	//fmt.Println(maxProfit([]int{1, 2}))
}

// maxProfit 找出最低点和最高点
func maxProfit(nums []int) int {
	n := len(nums)
	min := math.MaxInt32
	max := 0
	for i := 0; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i]-min > max {
			max = nums[i] - min
		}
	}
	return max
}

func maxProfit2(prices []int) (s int) {
	maxs := math.MinInt64
	mins := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		//记一个最小值
		mins = Min(mins, prices[i])
		//记一个最大值
		maxs = Max(maxs, prices[i]-mins)

	}
	//返回最大差值
	return maxs
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func Min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
