package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/17 10:47
 * @Description:杨辉三角
 */
func main() {
	fmt.Println(generate1(5))
}

// 观察规律可知道
func generate1(target int) [][]int {
	//分配好了内存空间
	//sum := make([][]int, target) //注意使用这种make初始化切片里面就是数据了，append时直接并接在其后面的了 fmpl([],[],[],[],[])
	var sum [][]int //从开始并接数据！fmpl([])
	fmt.Println(sum)
	//开始遍历
	for i := 0; i < target; i++ {
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
