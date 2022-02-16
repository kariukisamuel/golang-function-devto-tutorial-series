package main

import (
	"errors"
	"fmt"
)

func main() {
	answer, err := divide(2, 0)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%f", answer)
	}

}
//note we declare answer & err as return parameters
func divide(a, b float64) (answer float64, err error) {
	if a == 0 || b == 0 {
		err = errors.New("division by zero failed")
	}
	answer = a / b
	return
}
