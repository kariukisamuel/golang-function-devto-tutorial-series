package main

import "math"

func main() {
	p2 := powerOfTwo()
	value := p2()
	println(value)
	//print 4
	value = p2()
	println(value)
	//print 9
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
} 
