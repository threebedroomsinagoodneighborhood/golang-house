package rooms

import (
	"fmt"
	"house/assets/furniture"
)

type Bedroom struct {
	Bed     furniture.Bed
	Cabinet furniture.Cabinet
}

func CreateBedroom() Bedroom {
	fmt.Println("спальня: создание мебели")

	Bed := furniture.Bed{Furniture: furniture.Furniture{Type: "кровать", Material: "дерево", Color: "розовый", Length: 200, Width: 200, Height: 50}}
	Cabinet := furniture.Cabinet{Furniture: furniture.Furniture{Type: "шкафчик", Material: "дерево", Color: "коричневый", Length: 50, Width: 50, Height: 50}}

	Bed.Furniture.FurniturePrint()
	Cabinet.Furniture.FurniturePrint()

	bedroom := Bedroom{Bed: Bed, Cabinet: Cabinet}
	fmt.Println("спальня создана")
	return bedroom
}
