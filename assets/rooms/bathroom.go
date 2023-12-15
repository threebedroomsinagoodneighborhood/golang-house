package rooms

import (
	"fmt"
	"house/assets/furniture"
)

type Bathroom struct {
	Toilet  furniture.Toilet
	Bathtub furniture.Bathtub
}

func CreateBathroom() Bathroom {
	fmt.Println("ванная: создание мебели")

	Toilet := furniture.Toilet{Furniture: furniture.Furniture{Type: "туалет", Material: "керамика", Color: "белый", Length: 100, Width: 100, Height: 100}}
	Bathtub := furniture.Bathtub{Furniture: furniture.Furniture{Type: "ванная", Material: "керамика", Color: "белый", Length: 100, Width: 100, Height: 200}}

	Toilet.Furniture.FurniturePrint()
	Bathtub.Furniture.FurniturePrint()

	bathroom := Bathroom{Toilet: Toilet, Bathtub: Bathtub}
	fmt.Println("ванная создана")
	return bathroom
}
