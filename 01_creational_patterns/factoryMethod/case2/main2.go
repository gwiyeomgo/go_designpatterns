package main

import "fmt"

// 인터페이스 적용하기
type ShipFactory interface {
	order(name string) *Ship
}

func NewFactory(color string) ShipFactory {
	switch color {
	case "white":
		return WhiteShipFactory{}
	case "back":
		return BackShipFactory{}
	}
	return nil
}

type WhiteShipFactory struct{}

func (w WhiteShipFactory) order(name string) *Ship {
	return New(name, "white")
}

type BackShipFactory struct{}

func (b BackShipFactory) order(name string) *Ship {
	return New(name, "black")
}

func main() {
	whiteShipFactory := NewFactory("white")
	newShip := whiteShipFactory.order("whiteship1")
	fmt.Println(newShip.getName())

	//공장처럼 찍어낸다
	whiteShipFactory2 := NewFactory("white")
	newShip2 := whiteShipFactory2.order("whiteship2")
	fmt.Println(newShip2.getName())

	backShipFactory := NewFactory("back")
	newShip3 := backShipFactory.order("backship1")
	fmt.Println(newShip3.getName())
}
