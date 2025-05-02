package main

import "fmt"

// combination of OCP and Repository demo

type Product struct {
	name string
	color Color
	size Size
}


type Specification interface {
	IsSatisfied(p *Product) bool
}


func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{ "House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products:\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}






