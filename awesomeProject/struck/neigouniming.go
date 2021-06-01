package main

import (
	"encoding/json"
	"fmt"
)
// 车轮
type Wheel struct {
	Size int
}
// 车
type Car struct {
	Wheel
	// 引擎
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}
func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
	jsonData, _ := json.Marshal(c)
	//fmt.Println(jsonData)
    //newCar := struct {
	//	 Wheel
	//	 //Engine : struct{
	//	 //      Power int
	//	 //      Type  string
	//     //    }
	//}
	//json.Unmarshal(jsonData, &newCar)
	//fmt.Printf("%+v\n", newCar)

}
