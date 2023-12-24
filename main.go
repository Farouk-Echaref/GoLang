// this tells the compiler that our code should be compiled into an executable at the end 
package main

// importing fmt package to handle formating strings and printing messages
import (
	"fmt"
	"strings"
	"sort"
)

func getInitials(n string) (string, string){
	var initials []string

	splitted := strings.Split(n, " ")

	for _, val := range splitted{
		initials = append(initials, val[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

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
	arr2 := [3]string{"test", "test2", "test3"}
	fmt.Println(arr2)
	
	//slices (use array under the hood)
	slc := []int{1, 2, 3, 4, 5, 6}
	slc = append(slc, 69)
	fmt.Println(slc)

	//slice ranges
	rng := slc[0:3]
	fmt.Println(rng)

	//strings package
	phrase := "Testing strings package"
	fmt.Println(strings.Contains(phrase, "package"))
	//return a new one
	fmt.Println(strings.ReplaceAll(phrase, " ", "_"))
	fmt.Println(strings.Split(phrase, " "))

	//sort package
	nums := []int{4, 3, -1, 6, 0, 44, -4}
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	// loops
	x := 0
	for (x < 5){
		fmt.Println("x: ", x)
		x++
	}
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}
	strs := []string{"tes1", "tes2", "tes3", "tes4"}
	for index, value := range strs{
		fmt.Println(index, value)
	}
	for _, value := range strs{
		fmt.Print(value)
	}
	for i := 0; i < len(strs); i++{
		fmt.Println(strs[i])
	}

	fmt.Println(getInitials("Fuck"))

	MyPrint()
}