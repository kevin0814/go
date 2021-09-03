package main

import (
	"fmt"
	"github.com/taoshihan1991/imaptool/tools"
)

func main(){
	arr :=[]int{1,34,5,4,63,632,37,68}
	tools.BubbleSort(&arr)
	fmt.Println(arr)
	fmt.Println(GetMax(arr))
	fmt.Println(BubbleSort(arr))
}

//冒泡排序获取最大值
func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr[len(arr)-1]
}

//冒泡排序
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
