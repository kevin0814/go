package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
//sync.Once.Do(f func())能保证once只执行一次,这个sync.Once块只会执行一次。
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("v:", v, "---i:", i)
	}

	for i := 0; i < 10; i++ {

		go func(i int) {
			once.Do(onced)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
