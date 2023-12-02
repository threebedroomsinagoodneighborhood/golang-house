package rooms

import (
	"house/assets/furniture"
)

type Kitchen struct {
	Table  furniture.Table
	Chairs []furniture.Chair
	Fridge furniture.Fridge
	Stove  furniture.Stove
}

func CreateKitchen() Kitchen {

	var Chairs []furniture.Chair

	for i := 0; i < 4; i++ {
		Chairs = append(Chairs, furniture.Chair{Furniture: furniture.Furniture{Type: "стул", Color: "коричневый", Material: "Дерево", Length: 50, Width: 50, Height: 50}})
	}
	Table := furniture.Table{Furniture: furniture.Furniture{Type: "стол", Material: "дерево", Color: "коричневый", Length: 100, Width: 100, Height: 100}}
	Fridge := furniture.Fridge{Furniture: furniture.Furniture{Type: "холодильник", Material: "пластик", Color: "белый", Length: 100, Width: 100, Height: 200}}
	Stove := furniture.Stove{Furniture: furniture.Furniture{Type: "плита", Material: "стеклянная", Color: "белый", Length: 100, Width: 100, Height: 100}}

	kitchen := Kitchen{Table: Table, Chairs: Chairs, Fridge: Fridge, Stove: Stove}

	return kitchen
}
