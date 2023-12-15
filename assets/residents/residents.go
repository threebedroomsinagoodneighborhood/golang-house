package residents

import "fmt"

type Resident struct {
	Name string
	Age  int
}

type Family struct {
	Family []Resident
}

func (f Resident) FamilyPrint() {
	fmt.Println("\tроль: ", f.Name, "\n\tвозраст: ", f.Age, "\n")
}

func CreateFamily() Family {
	fmt.Println("жители: создание членов семьи")

	var family []Resident

	family = append(family, Resident{Name: "жена", Age: 42})
	family = append(family, Resident{Name: "муж", Age: 41})
	family = append(family, Resident{Name: "ребенок", Age: 13})
	for i := 0; i < 3; i++ {
		family[i].FamilyPrint()
	}
	fmt.Println("все жители созданы")
	return Family{Family: family}
}
