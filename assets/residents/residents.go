package residents

import "fmt"

type Resident struct {
	Name string
	Age  int
}

type Family struct {
	Family []Resident
}

func (f Resident) Print() {
	fmt.Print("имя ", f.Name, "\nвозраст: ", f.Age, "\n")
}

func CreateFamily() Family {

	var family []Resident

	family = append(family, Resident{Name: "жена", Age: 42})
	family = append(family, Resident{Name: "муж", Age: 41})
	family = append(family, Resident{Name: "ребенок", Age: 13})

	return Family{Family: family}
}
