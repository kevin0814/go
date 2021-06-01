package main

import ("fmt"

)
type newType struct {
	x int
	y int
}
//一个 struct 实现了某个接口里的所有方法，就叫做这个 struct 实现了该接口
func main() {
	fmt.Println("hello world")
	type Player struct{
		Name string
		HealthPoint int
		MagicPoint int
	}
	//创建指针类型的结构体
	tank := new(Player)
	tank.Name = "Canon"
	tank.HealthPoint = 300
	fmt.Println(tank)

    //取结构体的地址实例化
	//ins := &T{}  对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
    NR :=&Player{}
    NR.Name="张三"
    NR.HealthPoint=11
	NR.MagicPoint=100
	fmt.Println(NR)

	type People struct {
		name  string
		child *People
	}
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	fmt.Println(relation)

	type Interval struct {
		start  int
		end   int
	}
//	intr := Interval{0, 3}  //第一种赋值方式
	intr := Interval{end:5, start:1}
//	intr := Interval{end:5}
   fmt.Println(intr)
	//声明类型
	type Point3D struct { x, y, z float64 }
	type Line struct { p, q Point3D }
	  //结构体类型变量
	origin := Point3D{}  //  Point3D 是零值
	line := Line{origin, Point3D{y: -4, z: 12.3}}  //   line.q.x 是零值
	//这里 Point3D{}以及 Line{origin, Point3D{y: -4, z: 12.3}}都是结构体字面量。
	fmt.Println(line)
	//new(Type) 和 &Type{} 都是实例化构造体  &struct1{a, b, c} 是简写

}