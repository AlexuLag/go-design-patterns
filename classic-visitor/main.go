package main

import (
	"fmt"
)

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(de *DoubleExpression) {
	ee.result = de.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(ae *AdditionExpression) {
	ae.left.Accept(ee)
	x := ee.result
	ae.right.Accept(ee)
	y := ee.result
	ee.result = x + y
}
func main() {
	// 1+(2+3)
	e := &AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := &ExpressionEvaluator{}
	e.Accept(ee)
	fmt.Printf("%s = %g\n", ep, ee.result)

}
