package main

import (
	"fmt"
)

func main(){
	var mapLit  map [string]int
	var mapAssigned map [string]int
	mapLit  = map [string]int{"one":1,"two":2}
	mapCreated := make(map[string]float32)
	//mapCreated := make(map[string]float)等价于mapCreated := map[string]float{}
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	//map 可以设置容量大小，也可以不用
	//make(map[keytype]valuetype, cap)
	// := 等价于先声明变量，然后给变量赋值
	scene  :=make(map[string]int)
	scene["kev"] =1
	scene["vvv"] =3
	scene["wew"] =5
       //11
	for k,v :=range scene{
		fmt.Println(k,v)
	}

	slice := []int{10, 20, 30, 40}
	// 从第三个元素开始迭代每个元素
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
	scenes := make(map[string]int)
	// 准备map数据 
	scenes["route"] = 66
	scenes["brazil"] = 4
	scenes["china"] = 960
	delete(scenes, "brazil") 
	for k, v := range scenes {
		fmt.Println(k, v)
	}

}
