package main

import (
	"fmt"
	"sort"
	time "time"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/12 12:17
 * @Description:
 */
func main1() {
	m := 8 >> 1
	fmt.Println(m)
	m = 4 >> 1
	fmt.Println(m)
}
func main2() {
	c := make(chan int)
	v := <-c
	fmt.Println(v)
}

func main3() {

	ch := make(chan int, 10)
	defer close(ch) // 添加关闭channel的语句

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	for {
		select {
		case x := <-ch:
			println(x)
			time.Sleep(1 * time.Second)
		case <-time.After(1 * time.Second):
			println(time.Now().Unix())
		}
	}
}

func main4() {
	// 初始化 Map
	scores := map[string]int{
		"Tom":   90,
		"Jerry": 80,
		"Bob":   85,
		"Mike":  92,
	}

	// 将 Map 中的键复制到一个新的切片中
	keys := make([]string, 0, len(scores))
	for key := range scores {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Slice(keys, func(i, j int) bool {
		return scores[keys[i]] < scores[keys[j]]
	})

	// 打印排序后的结果
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, scores[key])
	}
}
func main5() {
	var oldSlice = []int64{1, 2, 3, 4, 5} // len:5,capacity:5
	var newSlice = oldSlice[1:3]          // len:2,capacity:4   (已经使用了两个位置，所以还空两位置可以append)
	fmt.Printf("%p\n", oldSlice)          //0xc420098000
	fmt.Printf("%p\n", newSlice)          //0xc420098008 可以看到newSlice的地址指向的是array[1]的地址，即他们底层使用的还是一个数组
	fmt.Printf("%v\n", oldSlice)          //[1 2 3 4 5]
	fmt.Printf("%v\n", newSlice)          //[2 3]

	newSlice[1] = 9              //更改后oldSlice、newSlice都改变了
	fmt.Printf("%v\n", oldSlice) // [1 2 9 4 5]
	fmt.Printf("%v\n", newSlice) // [2 9]

	newSlice = append(newSlice, 11, 12) //append 操作之后，oldSlice的len和capacity不变,newSlice的len变为4，capacity：4。因为这是对newSlice的操作
	fmt.Printf("%v\n", oldSlice)        //[1 2 9 11 12] //注意对newSlice做append操作之后，oldSlice[3],oldSlice[4]的值也发生了改变
	fmt.Printf("%v\n", newSlice)        //[2 9 11 12]

	newSlice = append(newSlice, 13, 14) // 因为newSlice的len已经等于capacity，所以再次append就会超过capacity值，
	// 此时，append函数内部会创建一个新的底层数组（是一个扩容过的数组），并将oldSlice指向的底层数组拷贝过去，然后在追加新的值。
	fmt.Printf("%p\n", oldSlice) //0xc420098000
	fmt.Printf("%p\n", newSlice) //0xc4200a0000
	fmt.Printf("%v\n", oldSlice) //[1 2 9 11 12]
	fmt.Printf("%v\n", newSlice) //[2 9 11 12 13 14]  它俩已经不再是指向同一个底层数组了
}
func main() {
	str := "nihao"
	fmt.Println(str[1:3])
}
