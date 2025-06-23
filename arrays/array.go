
// arrays size is fix so we can't add or remove elements

package main

import "fmt"

func main(){
	//int
	var numbers [5]int = [5]int{1,2,3,4,5}

	fmt.Println("numbers : ", numbers)

	fmt.Println("length of numbers : ", len(numbers))

	//string
	var names [3]string

	fmt.Println("names : ", names)

	names[0] = "John"
	
	fmt.Println("names : ", names)

	//bool
	var vals [3]bool

	fmt.Println("vals : ", vals)


	//short hand notation
	arr := [3]int{1,3,78}
	fmt.Println("arr : ", arr)

	//2d arrays

	matrix := [2][3]int{ {5,3,6},{6,4,3} }

	fmt.Println("matrix : ", matrix)
}