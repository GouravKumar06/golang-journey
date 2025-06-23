package main

import "fmt"


func main(){
	//defer

	for i := range 5  {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}