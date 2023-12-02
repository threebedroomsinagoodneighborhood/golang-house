package rooms

import "house/assets/furniture"

type Bedroom struct {
	Bed     furniture.Bed
	Cabinet furniture.Cabinet
}

func CreateBedroom() Bedroom {

	Bed := furniture.Bed{Furniture: furniture.Furniture{Type: "кровать", Material: "дерево", Color: "розовый", Length: 200, Width: 200, Height: 50}}
	Cabinet := furniture.Cabinet{Furniture: furniture.Furniture{Type: "шкафчик", Material: "дерево", Color: "коричневый", Length: 50, Width: 50, Height: 50}}

	bedroom := Bedroom{Bed: Bed, Cabinet: Cabinet}

	return bedroom
}
