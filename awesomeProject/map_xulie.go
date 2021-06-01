package main

import (
	"encoding/json"
	"fmt"
	"time"
)
func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	fmt.Println(sum)
}
func main() {

	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"]  = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age"      : 30,
		"hobby"    : []string{"读书","爬山"},
	}
	fmt.Println("map data :", res)
    test()
	//序列化
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data :", string(jsons))

	//反序列化
	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json marshal error:", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data :", res2)

	//编辑 删除
	person := map[int]string{
		1 : "Tom",
		2 : "Aaron",
		3 : "John",
	}
	fmt.Println("data :",person)

	delete(person, 2)
	fmt.Println("data :",person)

	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("data :",person)
}