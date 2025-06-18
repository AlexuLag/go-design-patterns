package main

import (
	"fmt"
)

type Expression interface {
	Accept(ev ExpressionVisitor)
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
