package main

import (
	"fmt"
)

type Human struct {
	name   string // 姓名
	Gender string // 性别
	Age    int    // 年龄
	string        // 匿名字段
}

type Student struct {
	Human     // 匿名字段
	Room  int // 教室
	int       // 匿名字段
}

func main() {
	//使用new方式
	stu := new(Student)//&Student{} 等价
	stu.Room = 102
	stu.Human.name = "Titan"
	stu.Gender = "男"
	stu.Human.Age = 14
	stu.Human.string = "Student"
// 内嵌结构体的字段，可以逐层选择来使用，如stu.Human.name。如果外层结构体中没有同名的name字段，也可以直接选择使用，如stu.name。
	fmt.Println("stu is:", stu)
	fmt.Printf("Student.Room is: %d\n", stu.Room)
	fmt.Printf("Student.int is: %d\n", stu.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stu.name) //  (*stu).name
	fmt.Printf("Student.Human.Gender is: %s\n", stu.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stu.Age)
	fmt.Printf("Student.Human.string is: %s\n", stu.string)
	//通过对结构体使用new(T)，struct{filed:value}两种方式来声明初始化，分别可以得到*T指针变量，和T值变量。
	// 使用结构体字面量赋值
	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}

	fmt.Println("stud is:", stud)
	fmt.Printf("Student.Room is: %d\n", stud.Room)
	fmt.Printf("Student.int is: %d\n", stud.int) // 初始化时已自动给予零值：0
	fmt.Printf("Student.Human.name is: %s\n", stud.Human.name)
	fmt.Printf("Student.Human.Gender is: %s\n", stud.Human.Gender)
	fmt.Printf("Student.Human.Age is: %d\n", stud.Human.Age)
	fmt.Printf("Student.Human.string is: %s\n", stud.Human.string)
}
