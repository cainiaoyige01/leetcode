package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: Niuzai
 * @Date: 2023/4/12 22:51
 * @Description:
 */
func main() {
	var wg sync.WaitGroup
	foo := make(chan int)
	bar := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(&bar)
		select {
		case v := <-bar:
			foo <- v
			fmt.Println(foo)
		default:
			println("default")
		}
		close(foo)
		close(bar)
	}()
	wg.Wait()
}
