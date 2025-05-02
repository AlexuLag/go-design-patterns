package main 

type Size int

const (
	small Size = iota
	medium
	large
)

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}
