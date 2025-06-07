package main

const age = 30
var email string = "gourav@gmail.com"
func main(){
	//constants
	const name string = "John Doe"
	println("name", name)

	println("age", age)

	println("email", email)

	email = "raj@gmail.com"
	println("email changed", email)

	//multiple constants
	const(
		port = 8080
		host = "localhost"
	) 

	println("port", port)
	println("host", host)

}