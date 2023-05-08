package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/4/12 0:03
  - @Description:编译报错 cannot assign to struct field m["foo"].x in map。错误原因：

对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。
对此 我们有两个方法来解决
1.采用临时变量来接受
2.采用指针类型
*/
type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}

func main() {
	//m["foo"].x = 4
	//采用第一种方法
	temp := m["foo"]
	temp.x = 4
	m["foo"] = temp
	fmt.Println(m["foo"].x)
}

// main1 第二种方法
func main1() {
	m1 := map[string]*Math{
		"foo": &Math{2, 3},
	}
	m1["foo"].x = 4
	fmt.Println(m["foo"].x)
}
