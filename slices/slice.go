package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic array
func main(){
    // uninitialized slice is nil

	var slice []int

	fmt.Println("slice", slice)  // []

	fmt.Println("length of slice", len(slice))   // 0 

	fmt.Println("equal to nil", slice == nil)  //true 

	// make   method
    var values = make([]int, 2,5)   // make([]type, length, capacity)   [ capacity is optional ]

	fmt.Println("values: ", values)
 
	//capacity -> maximum number of elements can fit [ automatically resizeable]

	fmt.Println("capacity: ", cap(values))

	// append [ add new element at the end of the slice]
	values = append(values, 3)

	fmt.Println("values: ", values)

	//short hand notation
	nums := []int{}

	fmt.Println("nums: ", nums)
	fmt.Println("length of nums: ", len(nums))
	fmt.Println("capacity of nums: ", cap(nums))


	// copy function   [ copy src, dst ]
	var numbers = make([]int, 0, 5)
	numbers = append(numbers, 1,2,3,4,5)
	var numbers2 = make([]int, len(numbers))

	fmt.Println("numbers: ", numbers)
	fmt.Println("numbers2: ", numbers2)

	copy(numbers2, numbers)  // copy numbers to numbers2

    fmt.Println("numbers2 after copy: ", numbers2)


	//slice operator

	var numbers3 = []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println("slicing values : ", numbers3[1:4])

	// slice package -> https://pkg.go.dev/slices

	var val1 = []int{1,2,3,4,5,6,7,8,9,10}
	var val2 = []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(slices.Equal(val1, val2))

	// 2 d slices
	var matrix = [][]int{{1,2,3},{4,5,6},{7,8,9}}

	fmt.Println("matrix: ", matrix)
}

