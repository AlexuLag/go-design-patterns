package main

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
