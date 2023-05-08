package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/15 9:40
 * @Description:加一
 */
func main() {
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{0}))
}

// plusOne 关键点是如果数字是9了 到10就要溢出了 这要做处理
func plusOne(nums []int) []int {
	n := len(nums)
	if nums[n-1]+1 < 10 {
		nums[n-1] = nums[n-1] + 1
	} else {
		for i := n - 1; i >= 0; i-- {
			if nums[i]+1 < 10 {
				nums[i] = nums[i] + 1
				break
			}
			nums[i] = 0
			if i == 0 {
				nums = append([]int{1}, nums...)
			}
		}
	}
	return nums
}

// plusOne2 对上述方法进行空间优化
func plusOne2(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] != 9 {
			nums[i]++
			for j := i + 1; j < n; j++ {
				nums[j] = 0
			}
			return nums
		}
	}
	//如果全是9999
	nums = make([]int, n+1)
	nums[0] = 1
	return nums

}
