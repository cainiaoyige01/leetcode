package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/12 10:28
 * @Description:旋转图片
 */
func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	matrix = rotate(matrix)
	fmt.Println(matrix)
}

// rotate 观察法！先把nums[i]与nums[n-1]等元素交换 采用对角线交换
func rotate(nums [][]int) [][]int {
	n := len(nums)
	//matrix := make([][]int, n)
	//muns2 := make([]int, n)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	fmt.Println(nums)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			nums[i][j], nums[j][i] = nums[j][i], nums[i][j]
		}

	}
	return nums
}

// rotate 使用辅助数组
func rotate1(nums [][]int) [][]int {
	n := len(nums)
	temp := make([][]int, n)
	//再分配[]int的存储空间
	for i := range temp {
		temp[i] = make([]int, n)
	}
	for i, row := range nums {
		for j, v := range row {
			temp[j][n-1-i] = v
		}
	}
	return temp
}
