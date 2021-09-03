package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func hellos(i int){
  defer  wg.Done() // goroutine结束就登记-1
  fmt.Println("hello goroutine !",i)
}
func main(){
	for i:=1; i<10;i++{
		wg.Add(1)// 启动一个goroutine就登记+1
		go hellos(i)
	}
	wg.Wait()// 等待所有登记的goroutine都结束
}
