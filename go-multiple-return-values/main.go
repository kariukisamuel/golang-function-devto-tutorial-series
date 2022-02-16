package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	answer, err := divide(2, 0)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%f", answer)
	}

}

func divide(a, b float64) (float64, error) {
	if a == 0 || b == 0 {
		return math.NaN(), errors.New("division by zero failed")
	}
	return a / b, nil
}
