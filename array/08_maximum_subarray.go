package main

import (
	"fmt"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/12 11:19
 * @Description:最大子序和
 */
func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}

// maxSubArray 动态规划 状态进行转移
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := nums[0]
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		dp[i] = max1(dp[i-1]+nums[i], nums[i])
		res = max1(dp[i], res)
	}
	return res
}
func max1(x, y int) int {
	if x > y {
		return x
	}
	return y
}
