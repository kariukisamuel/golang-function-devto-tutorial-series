package main

import "fmt"

func main() {
	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total := sum(12, 13, 14, 15.5)
	total2 := sum2(numbers...)
	fmt.Printf("Total sum %f", total)
	fmt.Printf("Total2 sum %f", total2)
}

func sum(values ...float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func sum2(values ...float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total

}
