package rooms

import (
	"house/assets/furniture"
)

type table struct {
	Furniture furniture.Furniture
}

type chair struct {
	Furniture furniture.Furniture
}
type cabinet struct {
	Furniture furniture.Furniture
}

type Living_room struct {
	Table   table
	Chairs  []chair
	Cabinet cabinet
}

func CreateLiving_room() Living_room {

	var Chairs []chair

	for i := 0; i < 4; i++ {
		Chairs = append(Chairs, chair{Furniture: furniture.Furniture{Type: "стул", Color: "коричневый", Material: "дерево", Length: 50, Width: 50, Height: 50}})
	}
	Table := table{Furniture: furniture.Furniture{Type: "стол", Material: "дерево", Color: "коричневый", Length: 100, Width: 100, Height: 100}}
	Cabinet := cabinet{Furniture: furniture.Furniture{Type: "шкафчик", Material: "дерево", Color: "коричневый", Length: 50, Width: 50, Height: 50}}

	living_room := Living_room{Table: Table, Chairs: Chairs, Cabinet: Cabinet}

	return living_room
}
