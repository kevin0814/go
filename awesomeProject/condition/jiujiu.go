package main

import "fmt"

func  main(){
	//遍历
	for y :=1;y<=9;y++ {
		for x:=1;x<=y;x++{
			fmt.Printf("%d*%d=%d ",x,y,x*y)
		}
		fmt.Println()
	}

}