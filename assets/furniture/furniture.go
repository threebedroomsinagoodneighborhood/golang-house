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

// кухня
type Table struct {
	Furniture Furniture
}

type Chair struct {
	Furniture Furniture
}
type Cabinet struct {
	Furniture Furniture
}

type Fridge struct {
	Furniture Furniture
}
type Stove struct {
	Furniture Furniture
}

// ванная
type Toilet struct {
	Furniture Furniture
}

type Bathtub struct {
	Furniture Furniture
}

// спальня
type Bed struct {
	Furniture Furniture
}
