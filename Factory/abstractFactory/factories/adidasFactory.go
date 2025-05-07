package factories

import (
	"abstractFactory/shirts"
	"abstractFactory/shoes"
)

type AdidasFactory struct {
}

func (a *AdidasFactory) MakeShoe() shoes.IShoe {
	return &shoes.AdidasShoe{
		Shoe: shoes.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (a *AdidasFactory) MakeShirt() shirts.IShirt {
	return &shirts.AdidasShirt{
		Shirt: shirts.Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
