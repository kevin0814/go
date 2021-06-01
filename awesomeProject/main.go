package main

import (
	"fmt"
	"net/http"
	"time"
)
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func sayHello(w http.ResponseWriter, r *http.Request){
      _,_ =fmt.Fprintln(w,"hello golang") //
}
func main() {
	http.HandleFunc("/hello",sayHello)
    err :=http.ListenAndServe(":9090",nil)

    if err !=nil{
    	fmt.Printf("http server fails")
	}
}
