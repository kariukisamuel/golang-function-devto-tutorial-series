package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

// function format is func _name(param1 param_type, param1 param_type) type_of_value_to_be_returned{}
func add(a, b int) int {
	return a + b
}
