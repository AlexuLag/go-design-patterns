package main

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &p)
		}
	}
	return result
}
