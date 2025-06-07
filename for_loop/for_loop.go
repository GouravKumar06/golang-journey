// for loop is the only loop in golang language
package main

import "fmt"

func main(){
	//while loop via for loop
    // i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//classic for loop
	for i := 0; i <= 2; i++  {
		fmt.Println("hello brother: ",i);
	}

	//range in loop
	for i := range 5 {
		fmt.Println(i)
	}
}