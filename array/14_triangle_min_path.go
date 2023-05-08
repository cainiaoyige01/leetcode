package main

import (
	"fmt"
	"math"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/17 21:02
 * @Description:三角形最小路径和
 */
func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},

		// {-1},
		// {2, 3},
		// {1, -1, -1},
	}
	fmt.Println(minimumTotal(triangle)) // 2 3 5 1
}

// 优化了数组！
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n)
	f[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + triangle[i][i]
		//返回来遍历时一大 【妙点】
		for j := i - 1; j > 0; j-- {
			f[i] = min2(f[j-1], f[j]) + triangle[i][j]
		}
		f[0] = f[0] + triangle[i][0]
	}
	sum := math.MaxInt32
	for i := range f {
		sum = min2(sum, f[i])
	}
	return sum
}

// minimumTotal 挺难做的！主要是边界问题要理解好！就是两个点最短路径 把一些公式列出来就好了！
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	total := make([][]int, n)
	for i := range total {
		total[i] = make([]int, i+1)
	}
	total[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		//第一个边界
		total[i][0] = total[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			total[i][j] = min2(total[i-1][j-1], total[i-1][j]) + triangle[i][j]
		}
		total[i][i] = total[i-1][i-1] + triangle[i][i]
	}
	var sum = math.MaxInt32
	for i := range total[n-1] {
		if total[n-1][i] < sum {
			sum = total[n-1][i]
		} else {
			continue
		}
	}
	return sum
}
func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
