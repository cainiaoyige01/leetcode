package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/14 19:53
  - @Description:不同路径
    一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
    机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
    问总共有多少条不同的路径？
*/
func main() {
	fmt.Println(uniquePaths3(3, 3)) // 7列 3行
}

// uniquePaths1 根据数学的规律可以知道是一个排列组合问题C(n,m) n表示向下走的步数 m表示总的步数
func uniquePaths1(row int, column int) int {
	//总步数至少也要一步
	sum := 1
	j := 1
	for i := (row + column - 2); i >= column; i-- {
		sum *= i
		sum /= j
		j++
	}
	return sum
}

// uniquePaths 使用动态规划 使用二维数组 规律：路径数等于nums[i-1][j]+nums[i][j-1]种
func uniquePaths2(row int, column int) int {
	//初始化
	nums := make([][]int, row+1)
	for i := range nums {
		nums[i] = make([]int, column+1)
	}
	fmt.Println(nums[1][1], nums[1][2])
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			if j == 1 && i == 1 {
				nums[i][j] = 1
			} else {
				nums[i][j] = nums[i-1][j] + nums[i][j-1]
			}
		}
	}
	return nums[row][column]
}

// uniquePaths 进行空间上的优化
func uniquePaths3(row int, column int) int {
	//初始化
	nums := make([]int, column+1)
	nums[0] = 0
	nums[1] = 1
	for i := 1; i <= row; i++ {
		for j := 1; j <= column; j++ {
			nums[j] = nums[j] + nums[j-1]
		}
	}
	return nums[column]
}
