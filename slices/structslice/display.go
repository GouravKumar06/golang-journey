package structslice

import "fmt"


func Display(){
	// slices from struct
	structSlice := []struct{
		name string
		age int
		admin bool
	}{
		{"John Doe", 30,false},
		{"Ramesh", 20,false},
		{"Suresh", 25,true},
		{"Raj", 30,false},
		{"Vijay", 40,false},
		{"bharti", 23,true},
	}

	fmt.Println("From Display():", structSlice)
}