package main

import "fmt"

func main() {
	//anonymous function are self invoking
	func() {
		fmt.Println("I am an anonymous function")
	}()
}
