package after

import (
	"fmt"
)

//https://refactoring.guru/ko/design-patterns/factory-method/go/example

func main() {
	factory := WhiteShipFactory{}
	whiteShip := factory.OrderShip("WhiteShip", "keesun@mail.com")
	fmt.Println(whiteShip.color)

	factory2 := BlackShipFactory{}
	blackShip := factory2.OrderShip("BlackShip", "keesun2@mail.com")
	fmt.Println(blackShip.color)
}
