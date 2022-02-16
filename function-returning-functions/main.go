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
	addExpr := mathExpression(AddExpr)
	println(addExpr(3, 2))
}
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
