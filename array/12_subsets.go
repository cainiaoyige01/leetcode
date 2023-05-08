package main

import "fmt"

/**
 * @Author: Niuzai
 * @Date: 2023/4/16 10:23
 * @Description:子集
 */
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}

// subsets 从后面往前遍历 每一遇到数就与所有的集合进行并接 然后再并接在大的集合中去
func subsets(nums []int) [][]int {
	n := len(nums)
	//先存[]值进去
	seqs := make([][]int, 1)
	//然后存3进去
	seqs = append(seqs, []int{nums[n-1]})
	//从后面遍历到前面
	for i := n - 2; i >= 0; i-- {
		for _, seq := range seqs {
			//并接进去 seq可能包含由多个值
			temp := append([]int{nums[i]}, seq...)
			seqs = append(seqs, temp)
		}
	}
	return seqs
}
