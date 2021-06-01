package main

import "fmt"

func main(){
	//fmt.PrintIn(nil==nil)
	fmt.Printf("%T\n",nil) //go 中空变量的表示
	//print(nil)

	var arr []int
	var num *int
	fmt.Printf("%p\n",arr)
	fmt.Printf("%p",num)
}
