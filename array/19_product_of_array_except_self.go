package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/22 15:23
 * @Description:除了自身以外数组的乘积
 */
func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4})) // [24 12 8 6]
	fmt.Println(productExceptSelf([]int{2, 4, 5, 7})) // [140 70 56 40]
}

// productExceptSelf 使用左子数的乘积与右子数的乘积 是一种很妙的做法
func productExceptSelf(nums []int) []int {
	n := len(nums)
	//定义一个数组
	result := make([]int, n)
	//先从左到右
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	//从右到左
	temp := nums[n-1]
	nums[n-1] = 1
	for i := n - 1 - 1; i >= 0; i-- {
		nums[i], temp = temp*nums[i+1], nums[i]
	}
	//遍历一遍就可以得到结果了
	for i := 0; i < n; i++ {
		nums[i] = nums[i] * result[i]
	}
	return nums
}
