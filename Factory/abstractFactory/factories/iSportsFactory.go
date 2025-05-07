package factories

import (
	"abstractFactory/shirts"
	"abstractFactory/shoes"
	"fmt"
)

type ISportsFactory interface {
	MakeShoe() shoes.IShoe
	MakeShirt() shirts.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &AdidasFactory{}, nil
	}

	if brand == "nike" {
		return &NikeFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
