// this tells the compiler that our code should be compiled into an executable at the end 
package main

// importing fmt package to handle formating strings and printing messages
import "fmt"

// main function that gets called when the program starts
func main(){

	// use capital(Println) if you want to export methods 
	fmt.Println("Testing go")

	//strings
	//explicit declaration
	var str1 string = "val"
	//implicit
	var str2 = "val2"
	//default value is empty string
	var str3 string

	fmt.Println(str1, str2, str3)

	//another syntax of declaring variables(should be inside a function)
	str4 := "val3"
	str3 = str4

	fmt.Println(str1, str2, str3, str4)
	
	//int & float 
	var one int = 1
	var two = 2
	three := 3
	var four float64 = 3.14

	fmt.Println(one, two, three, four)

	//Print method(without newline)
	fmt.Print("Testing ")
	fmt.Print("Line \n")

	//arrays
	// var arr [3]int = [3]int{1, 2, 3}
	var arr = [3]int{1, 2, 3}

	//print the array and it's length
	fmt.Println(arr, len(arr))
}