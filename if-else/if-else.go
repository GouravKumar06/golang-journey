// go does not have ternary operator
package main

import "fmt"

func main() {
	var age int = 18
	if age == 18 {
		fmt.Println("You can vote: ", age)
	} else {
		fmt.Println("You can't vote: ",age)
	}

	// 2. way 
	if hasPermission := false; hasPermission{
		fmt.Println("You can vote")
	}else{
		fmt.Println("You can't vote")
	}

}