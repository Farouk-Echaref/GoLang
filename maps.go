package main

import "fmt"

func main(){

	//declaring map
	maps := map[string]float64{
		"one": 1,
		"two": 2,
		"three": 3,
	}

	fmt.Println(maps)
	fmt.Println(maps["two"])

	//looping through map
	for k, v := range maps{
		fmt.Println(k, " = ", v)
	}
	maps2 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	for k, v := range maps2{
		fmt.Println(k, " = ", v)
	}
}