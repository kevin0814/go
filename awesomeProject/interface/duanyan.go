package main

import ("fmt")

func main (){
	var x interface{}
	x =10
	value, ok :=x.(int)
	fmt.Println(value,",",ok)

	var  a int
	a =10
	getType(a)
}

func getType(a interface{}){
	switch a.(type) {
	case  int :
		fmt.Println("int")
	case  string :
		fmt.Println("string")
	case  float64 :
		fmt.Println("float64")
	default:
		fmt.Println("unknown type")
	}

}
