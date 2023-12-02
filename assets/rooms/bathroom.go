package rooms

import "house/assets/furniture"

type Bathroom struct {
	Toilet  furniture.Toilet
	Bathtub furniture.Bathtub
}

func CreateBathroom() Bathroom {

	Toilet := furniture.Toilet{Furniture: furniture.Furniture{Type: "туалет", Material: "керамика", Color: "белый", Length: 100, Width: 100, Height: 100}}
	Bathtub := furniture.Bathtub{Furniture: furniture.Furniture{Type: "ванная", Material: "керамика", Color: "белый", Length: 100, Width: 100, Height: 200}}

	bathroom := Bathroom{Toilet: Toilet, Bathtub: Bathtub}

	return bathroom
}
