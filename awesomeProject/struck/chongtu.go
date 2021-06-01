package main

import (
	"fmt"
)
type A struct {
	a int
	b string
}
type B struct {
	b int
}
type C struct {
	A
	B
}
func main() {
	//jsonData := getJsonData()
	c := &C{}
	c.A.a = 1
	c.A.b = "222"
	fmt.Println(c)
	//screenAndTouch := struct {
	//	A
	//	B
	//	   }{}
	//jsonData := json.Marshal(c,&screenAndTouch)
	//fmt.Println("%+v\n",jsonData)
}