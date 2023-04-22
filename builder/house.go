package main

import "fmt"

type House struct {
	Doors   int
	Rooms   int
	Windows int
	Floor   int
}

type HouseInterface interface {
	SetDoors(int)
	SetRooms(int)
	SetWindows(int)
	SetFloor(int)
	PrintData()
}

func NewHouse() HouseInterface {
	return &House{}
}

func (h *House) SetDoors(d int) {
	h.Doors = d
}

func (h *House) SetRooms(r int) {
	h.Rooms = r
}

func (h *House) SetWindows(w int) {
	h.Windows = w
}

func (h *House) SetFloor(f int) {
	h.Floor = f
}

func (h *House) PrintData() {
	fmt.Printf(
		"Doors : %d\nFloor : %d\nRooms : %d\nWindows : %d\n",
		h.Doors, h.Floor, h.Rooms, h.Windows,
	)
}

func main() {
	house := NewHouse()
	house.SetDoors(5)
	house.SetFloor(12)
	house.SetRooms(8)
	house.SetWindows(6)
	house.PrintData()
}
