package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/15 8:44
 * @Description:最小路径和
 */
func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

// minPathSum 是利用不同路径的类似方法 求到某个点的最小距离
func minPathSum(nums [][]int) int {
	if len(nums) == 0 || len(nums[0]) < 0 {
		return 0
	}
	//定义一个全新的切片来存储 新的数据 这个点类似于螺旋矩阵
	minNums := make([][]int, len(nums))
	//分配空间了
	for i := range minNums {
		minNums[i] = make([]int, len(nums[i]))
	}
	//先分配第一行 以及第一列
	minNums[0][0] = nums[0][0]
	//减少每一内存开销创建
	row := len(nums)
	column := len(nums[0])
	for i := 1; i < row; i++ {
		minNums[i][0] = nums[i][0] + minNums[i-1][0]
	}
	for i := 1; i < column; i++ {
		minNums[0][i] = nums[0][i] + minNums[0][i-1]
	}
	//求其余的点的思想是到这点的最小值加上它本身就是到这点的最小距离了
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			minNums[i][j] = min(minNums[i-1][j], minNums[i][j-1]) + nums[i][j]
		}
	}
	return minNums[row-1][column-1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
