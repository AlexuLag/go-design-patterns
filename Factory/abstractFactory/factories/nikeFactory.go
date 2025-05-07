package factories

import (
	"abstractFactory/shirts"
	"abstractFactory/shoes"
)

type NikeFactory struct {
}

func (n *NikeFactory) MakeShoe() shoes.IShoe {
	return &shoes.NikeShoe{
		Shoe: shoes.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *NikeFactory) MakeShirt() shirts.IShirt {
	return &shirts.NikeShirt{
		Shirt: shirts.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
