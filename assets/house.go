package assets

import (
	"fmt"
	"house/assets/residents"
	"house/assets/rooms"
)

type Rooms struct {
	Living_room rooms.Living_room
	Bedroom     rooms.Bedroom
	Bathroom    rooms.Bathroom
	Kitchen     rooms.Kitchen
}

type House struct {
	Residents residents.Family
	Rooms     Rooms
	Length    float32
	Width     float32
	Height    float32
}

func CreateHouse() House {
	fmt.Println("создание дома начато:\n")
	fmt.Println("создание комнат\n")
	room_list := Rooms{
		Living_room: rooms.CreateLiving_room(),
		Bedroom:     rooms.CreateBedroom(),
		Bathroom:    rooms.CreateBathroom(),
		Kitchen:     rooms.CreateKitchen()}
	fmt.Println("\nвсе комнаты созданы\n")
	house := House{Rooms: room_list,
		Residents: residents.CreateFamily(),
		Length:    250, Width: 250, Height: 250,
	}
	fmt.Println("\nсоздание дома завершено.")
	return house
}
