package main

import "fmt"



type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//pointers and struct work together
	p := &v
	p.X = 1e9
	fmt.Println(v)

	display()
}
