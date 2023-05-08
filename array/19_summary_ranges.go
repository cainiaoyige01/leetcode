package main

import (
	"fmt"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/21 23:12
 * @Description:汇总区间
 */
func main() {
	fmt.Printf("%q\n", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0->2" "4->5" "7"]
}

func summaryRanges(nums []int) (sum []string) {
	n := len(nums)
	if n == 0 {
		return nil
	}
	slow, fast := 0, 0
	for fast < n-1 {
		if nums[fast]+1 == nums[fast+1] {
			fast++
		} else {
			if fast-slow == 0 {
				sprintf := fmt.Sprintf("%d", nums[slow])
				sum = append(sum, sprintf)
			} else {
				sprintf := fmt.Sprintf("%d->%d", nums[slow], nums[fast])
				sum = append(sum, sprintf)
			}
			fast++
			slow = fast
		}
	}
	if fast-slow == 0 {
		sprintf := fmt.Sprintf("%d", nums[slow])
		sum = append(sum, sprintf)
	} else {
		sprintf := fmt.Sprintf("%d->%d", nums[slow], nums[fast])
		sum = append(sum, sprintf)
	}
	return
}
