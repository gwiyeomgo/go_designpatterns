package main

import (
	"fmt"
)

type IShip interface {
	setColor(name string)
	setName(name string)
	getName() string
	getShip() Ship
}

type Ship struct {
	color string
	name  string
}

func (s *Ship) setColor(color string) {
	s.color = color
}
func (s *Ship) setName(name string) {
	s.name = name
}
func (s *Ship) getName() string {
	return s.name
}
func (s *Ship) getShip() Ship {
	return *s
}

type ShipFactory struct {
}

func (f ShipFactory) orderShip(name string, email string) Ship {
	// validate
	//new Ship
	/*	ship := Ship{}
		ship.setName(name)*/
	// coloring
	var ship Ship
	switch name {
	case "WhiteShip":
		ship = newWhiteShip().getShip()
	case "BlackShip":
		ship = newBlackShip().getShip()
	}

	// notify
	fmt.Printf("%s 로 email 전송 ,%s 를 만들었습니다.", email, ship.getName())

	return ship
}

//https://refactoring.guru/ko/design-patterns/factory-method/go/example

type WhiteShipFactory struct {
	ShipFactory
}
type WhiteShip struct {
	Ship
}

func newWhiteShip() IShip {
	return &WhiteShip{
		Ship: Ship{
			color: "white",
			name:  "WhiteShip",
		},
	}
}

type BlackShipFactory struct {
	ShipFactory
}
type BlackShip struct {
	Ship
}

func newBlackShip() IShip {
	return &BlackShip{
		Ship: Ship{
			color: "black",
			name:  "BlackShip",
		},
	}
}
func main() {
	factory := WhiteShipFactory{}
	whiteShip := factory.orderShip("WhiteShip", "keesun@mail.com")
	fmt.Println(whiteShip.color)

	factory2 := BlackShipFactory{}
	blackShip := factory2.orderShip("BlackShip", "keesun2@mail.com")
	fmt.Println(blackShip.color)
}
