package after

import "fmt"

type ShipFactory interface {
	OrderShip(name string, email string) Ship
}

//https://smyrman.medium.com/go2-interface-default-methods-bf487b49cf9c

func OrderShip(name string, email string) Ship {
	// validate
	//new Ship
	/*	ship := Ship{}
		ship.setName(name)*/
	// coloring
	var ship Ship
	switch name {
	case "WhiteShip":
		ship = NewWhiteShip().getShip()
	case "BlackShip":
		ship = NewBlackShip().getShip()
	}

	// notify
	fmt.Printf("%s 로 email 전송 ,%s 를 만들었습니다.", email, ship.getName())

	return ship
}

//new facotry

type WhiteShipFactory struct {
	ShipFactory
}
type WhiteShip struct {
	Ship
}

func NewWhiteShip() IShip {
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

func NewBlackShip() IShip {
	return &BlackShip{
		Ship: Ship{
			color: "black",
			name:  "BlackShip",
		},
	}
}
