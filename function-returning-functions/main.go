package main

import "github.com/kariukisamuel/function-returning-functions/simplemath"

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
	DivideExpr   = MathExpr("divide")
)

func main() {
	// addExpr := mathExpression(AddExpr)
	// // println(AddExpr)
	// // println(addExpr(3, 2))
	println(double(3, 2, mathExpression(SubtractExpr)))
}

//returning function from a function
func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	case DivideExpr:
		return simplemath.Divide
	default:
		return func(f, f2 float64) float64 {
			return 0
		}
	}
}

//passing function as a parameter
func double(x float64, y float64, mxpr func(float64, float64) float64) float64 {
	return 2 * mxpr(x, y)
}
