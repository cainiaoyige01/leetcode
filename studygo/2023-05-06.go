package main

import "fmt"

/*
*
  - @Author: Niuzai
  - @Date: 2023/5/6 12:08
  - @Description: 这相当于对map的元素可以直接取地址吗？
    即是无法对map的value或者key进行取值的
    如果通过其他 hack 的方式，例如 unsafe.Pointer 等获取到了 key 或 value 的地址，
    也不能长期持有，因为一旦发生扩容，key 和 value 的位置就会改变，之前保存的地址也就失效了
*/
func main() {
	m := make(map[string]int)
	fmt.Println(m["niuzai"])
}
