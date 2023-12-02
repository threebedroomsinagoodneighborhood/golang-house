package assets

import (
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
	rooms := Rooms{Living_room: rooms.CreateLiving_room(),
		Bedroom:  rooms.CreateBedroom(),
		Bathroom: rooms.CreateBathroom(),
		Kitchen:  rooms.CreateKitchen()}

	house := House{Rooms: rooms,
		Residents: residents.CreateFamily(),
		Length:    250, Width: 250, Height: 250,
	}
	return house
}
