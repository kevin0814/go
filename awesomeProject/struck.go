package main
import ("fmt"
    	"sync"
)
    // struct 结构体
	type Person struct {
		Name string
		Age int
	}

type person struct {
	name string
	city string
	age  int8
}

	func main() {
		p3 := &person{} //使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
		fmt.Printf("%T\n", p3)     //*main.person
		fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
		p3.name = "博客"
		p3.age = 30
		p3.city = "成都"
		fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}

		var p1 Person
		p1.Name = "Tom"
		p1.Age  = 30
		fmt.Println("p1 =", p1)

		var p2 = Person{Name:"Burke", Age:31}
		fmt.Println("p2 =", p2)

		//p3 := Person{Name:"Aaron", Age:32}
		//fmt.Println("p2 =", p3)

		//匿名结构体
		p4 := struct {
			Name string
			Age int
		} {Name:"匿名", Age:33}
		fmt.Println("p4 =", p4)

		//同步机构体
		var scene sync.Map
		// 将键值对保存到sync.Map
		scene.Store("greece", 97)
		scene.Store("london", 100)
		scene.Store("egypt", 200)
		// 从sync.Map中根据键取值
		fmt.Println(scene.Load("london"))
		// 根据键删除对应的键值对
		scene.Delete("london")
		// 遍历所有sync.Map中的键值对
		scene.Range(func(k, v interface{}) bool {
			fmt.Println("iterate:", k, v)
			return true
		})
	}

