package main

import "fmt"

func main() {
	//anonymous function are self invoking
	func() {
		fmt.Println("I am an anonymous function")
	}()
	//variables as function
	a := func(name string) string {
		return name
	}
	b := a("sam")
	println(b)

}
