package main

import (
	"fmt"
	"strconv"
)

/**
 * @Author: _niuzai
 * @Date:   2023/5/13 10:50
 * @Description:二进制求和
 */
func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

// addBinary3 写在力扣上的题解
func addBinary3(a string, b string) string {
	//首先获取各长度的进行比较便利
	lenA := len(a)
	lenB := len(b)
	//确定好最长的字符串 定下了便利的次数
	n := 0
	if lenA > lenB {
		n = lenA
	} else {
		n = lenB
	}
	//定义一些变量来接受溢出值
	caary := 0
	res := ""
	for i := 0; i < n; i++ {
		//判断每一次字符串是否已经结束了
		if i < lenA {
			//注意从字符串中取出的数据都是字符
			caary += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			caary += int(b[lenB-i-1] - '0')
		}
		//这里开始运算以及并接成字符串
		res = strconv.Itoa(caary%2) + res
		caary /= 2
	}
	//判断carry是否还有值 有的也只能是溢出的1
	if caary > 0 {
		res = "1" + res
	}
	return res
}

// addBinary
func addBinary(a string, b string) string {
	i1, _ := strconv.ParseInt(a, 10, 32)
	i2, _ := strconv.ParseInt(b, 10, 32)
	sum := i1 + i2
	var count int64 = 0
	var res []int64
	for sum != 0 {
		sum = sum + count
		temp := sum % 10
		res = append(res, temp%2)
		count = temp / 2
		sum = sum / 10
	}
	result := ""

	n := len(res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	for _, re := range res {
		result = fmt.Sprintf("%s%d", result, re)
		//result = result + "" + string(re)
	}
	if count != 0 {
		result = fmt.Sprintf("%d%s", count, result)
		//result = string(count) + "" + result
	}
	return result

}

// addBinary使用官方的方法
func addBinary2(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	//确定最大的方便接下来进行遍历
	n := Max(lenA, lenB)
	//定义一个变量来接受溢出的数据
	carry := 0
	ans := ""
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-1-i] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-1-i] - '0')
		}
		//数字转成字符串
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
