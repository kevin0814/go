package main

import "fmt"

func main(){
	//fmt.PrintIn(nil==nil)
	fmt.Printf("%T\n",nil)
	//print(nil)

	var arr []int
	var num *int
	fmt.Printf("%p\n",arr)
	fmt.Printf("%p",num)
}