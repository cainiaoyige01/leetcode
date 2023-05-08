package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/13 9:39
 * @Description:螺旋矩阵
 */
func main() {
	fmt.Println(spiralOrder([][]int{
		// {1, 2, 3, 4},
		// {5, 6, 7, 8},
		// {9, 10, 11, 12},

		// {1, 2},
		// {3, 4},

		{2, 3, 4},
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
		{14, 15, 16},
	}))
}

// spiralOrder 碰壁思想很精妙
func spiralOrder(nums [][]int) []int {
	//定义边界先
	var result []int
	if len(nums) == 0 {
		return result
	}
	//注意长宽
	left, up, right, down := 0, 0, len(nums[0])-1, len(nums)-1
	//定义一个坐标 分别代表行和列
	var x, y int
	for left <= right && up <= down {
		for y = left; y <= right && avoid(left, right, up, down); y++ {
			result = append(result, nums[x][y])
		}
		y--
		up++
		for x = up; x <= down && avoid(left, right, up, down); x++ {
			result = append(result, nums[x][y])
		}
		x--
		right--
		for y = right; y >= left && avoid(left, right, up, down); y-- {
			result = append(result, nums[x][y])
		}
		y++
		down--
		for x = down; x >= up && avoid(left, right, up, down); x-- {
			result = append(result, nums[x][y])
		}
		x++
		left++
	}
	return result
}
func avoid(left, right, up, down int) bool {
	return up <= down && left <= right
}
//spiralOrder1 还有简单一点的！进行了优化
func spiralOrder1(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	if n == 0 {
		return nil
	}
	top := 0
	bottom := m - 1
	left := 0
	right := n - 1
	var res []int
	for top <= bottom && left <= right {
		// 上
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// 右
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 下
		for i := right; i >= left && bottom >= top; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		// 左
		for i := bottom; i >= top && left <= right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
