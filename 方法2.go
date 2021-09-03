package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}
type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}
//该 method 属于 Circle 类型对象中的方法
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	emp1 := Employee {
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	//
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
