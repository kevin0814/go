package main

import (
	"fmt"
	"sync"
)

var wgs sync.WaitGroup

func hello() {
	defer wgs.Done() //减一 defer 延迟调用。一般是在return的时候调用
	fmt.Println("Hello Goroutine!")
}
func main() {
	wgs.Add(1) //加一 协程个数
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wgs.Wait()
}
