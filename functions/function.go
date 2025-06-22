
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x,y int) int {
	return x - y
}

func swap(x string, y string) (string,string) {
	return y,x
}

func split(sum int) (x,y float32) {
	x = float32(sum) * 4 / 9
	y = float32(float32(sum) - x)
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sub(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)	

	fmt.Println(split(17))
}
