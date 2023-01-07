package main

import "fmt"

//https://refactoring.guru/ko/design-patterns/factory-method/go/example

type Ship struct {
	name  string
	color string
}

func (s Ship) getName() {
	fmt.Printf(s.name)
}

func NewFactory(color string) func(name string) *Ship {
	return func(name string) *Ship {
		return &Ship{name: name, color: color}
	}
}

func main() {
	whiteShipFactory := NewFactory("white")
	newShip := whiteShipFactory("whiteship1")
	newShip.getName()

	//공장처럼 찍어낸다
	whiteShipFactory2 := NewFactory("white")
	newShip2 := whiteShipFactory2("whiteship2")
	newShip2.getName()

	backShipFactory := NewFactory("back")
	newShip3 := backShipFactory("backship1")
	newShip3.getName()
}
