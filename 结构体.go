package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)

	//结构体指针
	//指针指向一个结构体
	//也可以创建指向结构的指针。
	/*
		结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
			结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”
	*/
	//1.方法一
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 30
	p1.sex = "男"
	p1.address = "北京市"
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p1.name,p1.age,p1.sex,p1.address)
	fmt.Printf("%p,%T\n",&p1,p1)
	pn := p1
	fmt.Println(pn)
	fmt.Printf("%p,%T\n",&pn,pn)

	//2.方法二
	p2 := Person{}
	p2.name = "Ruby"
	p2.age = 28
	p2.sex= "女"
	p2.address = "上海市"
	fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p2.name,p2.age,p2.sex,p2.address)
	fmt.Printf("%p,%T\n",&p2,p2)

	//3.方法三
	p3 := Person{name :"如花",age :20,sex:"女",address:"杭州市"}
	fmt.Println(p3)
	p4 := Person{
		name:"隔壁老王",
		age : 40,
		sex :"男",
		address:"武汉市",
	}
	fmt.Println(p4)

	//4.方法四
	p5 := Person{"李小花",25,"女","成都"}
	fmt.Println(p5)
}

//1.定义结构体
type Person struct {
	name string
	age int
	sex string
	address string
}
