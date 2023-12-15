package furniture

import "fmt"

type Furniture struct {
	Type     string
	Color    string
	Material string
	Length   float32
	Width    float32
	Height   float32
}

func (f Furniture) FurniturePrint() {
	fmt.Println("\n\tтип мебели: ", f.Type, "\n\tцвет: ", f.Color, "\n\tматериал: ", f.Material, "\n\tдлина: ", f.Length, "\n\tширина: ", f.Width, "\n\tвысота: ", f.Height, "\n")
}

// общие
type Table struct { //стол
	Furniture Furniture
}

type Chair struct { //стул
	Furniture Furniture
}
type Cabinet struct { //шкафчик
	Furniture Furniture
}

// кухня
type Fridge struct { //холодильник
	Furniture Furniture
}
type Stove struct { //плита
	Furniture Furniture
}

// ванная
type Toilet struct { //туалет
	Furniture Furniture
}

type Bathtub struct { //ванная
	Furniture Furniture
}

// спальня
type Bed struct { //кровать
	Furniture Furniture
}
