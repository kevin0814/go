package main

import "fmt"

func main(){
	//使用make()创建map map 创建
   //var map_variable map[key_data_type]value_data_type
	/* 使用 make 函数 */
	//map_variable = make(map[key_data_type]value_data_type)
	rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	fmt.Printf("%v\n",rating)
	//如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	var countryCapitalMap map[string]string
	/* 创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}

	// delete 函數
	//delete(map, key) 函数用于删除集合的元素, 参数为 map 和其对应的 key。删除函数不返回任何值。
	/* 创建 map */
	countryCapitalMaps := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMaps {
		fmt.Println("Capital of",country,"is",countryCapitalMaps[country])
	}

	/* 删除元素 */
	delete(countryCapitalMaps,"France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素France后 map")

	/* 打印 map */
	for country := range countryCapitalMaps {
		fmt.Println("Capital of",country,"is",countryCapitalMaps[country])
	}
	//ok-idiom map[key]
	//但是当key如果不存在的时候，我们会得到该value值类型的默认值，比如string类型得到空字符串，int类型得到0。但是程序不会报错。
   //value, ok := map[key]
	m := make(map[string]int)
	m["a"] = 1
	x, ok := m["b"]
	fmt.Println(x, ok)
	x, ok = m["a"]
	fmt.Println(x, ok)

	//len(map)  可以得到map的长度
	// map是引用类型的 当将映射分配给一个新变量时 它们都指向相同的内部数据结构。因此，一个的变化会反映另一个
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)
	//map不能使用==操作符进行比较
}
