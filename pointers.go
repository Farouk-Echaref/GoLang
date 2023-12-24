package main

import "fmt"

func updateStr(str *string){
	*str = "newStr"
}

func main(){
	name := "testing"

	// m := &name

	var m *string = &name
	fmt.Println(*m)
	fmt.Println(name)
	
	// updateStr(&name)
	updateStr(m)

	fmt.Println(*m)
	fmt.Println(name)
}