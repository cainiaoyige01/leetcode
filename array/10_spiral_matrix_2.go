package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/14 19:31
 * @Description:螺旋矩阵||
 *给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * 关键点呀！从1开始升序 而且是n x n 的矩阵
 */
func main() {
	matrix := generateMatrix(3)
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

// generateMatrix 就和09_spiral_matrix一样的思想！这个是定义好矩阵 然后填数字
func generateMatrix(matrixLength int) [][]int {
	//首先初始化了一个矩阵
	nums := make([][]int, matrixLength)
	//分配切片的具体空间
	for i := range nums {
		nums[i] = make([]int, matrixLength)
	}
	//定义一个当前行与列 进行移动遍历 使用老办法碰壁思想
	r, c := 0, 0
	for i := 1; i <= matrixLength*matrixLength; {
		//首先从左往右移动
		for c < matrixLength && nums[r][c] == 0 {
			nums[r][c] = i
			c++
			i++
		}
		c--
		r++
		//从上往下
		for r < matrixLength && nums[r][c] == 0 {
			nums[r][c] = i
			r++
			i++
		}
		c--
		r-- //属于r==matrixLength 要回退一格
		//从右往左
		for c >= 0 && nums[r][c] == 0 {
			nums[r][c] = i
			c--
			i++
		}
		c++
		r--
		//从下往上
		for r >= 0 && nums[r][c] == 0 {
			nums[r][c] = i
			i++
			r--
		}
		r++
		c++
	}
	return nums
}
