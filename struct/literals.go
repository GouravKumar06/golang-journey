
package main

import "fmt"

type Nums struct {
	X, Y int
}

var (
	v1 = Nums{1, 2}  // has type Nums
	v2 = Nums{X: 1}  // Y:0 is implicit
	v3 = Nums{}      // X:0 and Y:0
	p  = &Nums{1, 2} // has type *Nums
)

func display() {
	fmt.Println(v1, p, v2, v3)
}
