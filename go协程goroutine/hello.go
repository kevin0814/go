package main

import (
	"fmt"
)
func hello(){
	fmt.Println("hello goroutine 6……")
}
func main(){
	go hello()
	fmt.Println("main go goroutine done!")
	//time.Sleep(time.Second)
}
