package main 

type Color int

const (
	red Color = iota
	green
	blue
)

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

