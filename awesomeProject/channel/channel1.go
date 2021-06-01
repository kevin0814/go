package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // 不使用带缓冲区的通道
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}

// 只能向chan里send数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ready ", i)  //管道沒有緩存區時，写入之后，就堵塞通道
		c <- i
		fmt.Println("send ", i)
	}
}

// 只能接收通道中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}