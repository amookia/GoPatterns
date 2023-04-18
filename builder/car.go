package main

import "fmt"

type Car struct {
	Color string
	Speed int
}

type CarInterface interface {
	SetColor(string)
	SetSpeed(int)
	printCar()
}

func NewCar() CarInterface {
	return &Car{}
}

func (c *Car) SetColor(color string) {
	c.Color = color
}

func (c *Car) SetSpeed(speed int) {
	c.Speed = speed
}

func (c *Car) printCar() {
	fmt.Println("color : ", c.Color)
	fmt.Println("speed : ", c.Speed)
}

func main() {
	c := NewCar()
	c.SetColor("Red")
	c.SetSpeed(100)
	c.printCar()
}
