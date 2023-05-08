package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/17 11:13
 * @Description:杨辉三角|| 要求空间复杂度为O（k）
 */
func main() {
	triangle := pascalsTriangle(3)
	fmt.Println(triangle[3])
}

func pascalsTriangle(target int) [][]int {
	//var result []int
	//分配好了内存空间
	//sum := make([][]int, target) //注意使用这种make初始化切片里面就是数据了，append时直接并接在其后面的了 fmpl([],[],[],[],[])
	//但是可以使用range 遍历 然后修改每每一个数组就好了
	var sum [][]int //从开始并接数据！fmpl([])
	//开始遍历
	for i := 0; i <= target; i++ {
		//分配第二层切片的空间
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 1; j < i; j++ {
			temp[j] = sum[i-1][j-1] + sum[i-1][j]
		}
		sum = append(sum, temp)
	}
	return sum
}

// pascalsTriangle 优化空间复杂度为O（k） 使用动态规划 关键点是公式nums[i]=nums[i-1]*(target-i+i)/i 时间复杂度也低
func pascalsTriangle1(target int) []int {
	nums := make([]int, target+1)
	nums[0] = 1
	for i := 1; i <= target; i++ {
		nums[i] = nums[i-1] * (target - i + 1) / i
	}
	return nums
}

// 优化空间复杂度 但是时间复杂度为n^2
func pascalsTriangle2(target int) []int {
	nums := make([]int, target+1)
	nums[0] = 1
	for i := 1; i <= target; i++ {
		for j := i; j >= 1; j-- {
			nums[j] = nums[j] + nums[j-1]
		}
	}
	return nums
}

// 使用双数组
func pascalsTriangle3(target int) []int {
	var cur, temp []int
	for i := 0; i <= target; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = temp[j-1] + temp[j]
		}
		temp = cur
	}
	return temp
}
