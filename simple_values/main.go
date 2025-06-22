package main

var name string = "John Doe"
var age int = 30
var isPasswordMatched bool = true

func main() {
	println(name, age, isPasswordMatched)
	//simple values
	var i int = 10
	
	var isPasswordMatched bool = true
	var s string = "Hello World"
	println(i, isPasswordMatched, s)
	
	display()
}

func display() {
	println("display: ==> ",name, age, isPasswordMatched)
}	




