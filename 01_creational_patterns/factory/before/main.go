package main

import "fmt"

// 어떤 요구사항으로 계속 변경됨
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

type ShipFactory struct {
}

func (f ShipFactory) orderShip(name string, email string) Ship {
	// validate
	//new Ship
	ship := Ship{}
	ship.setName(name)

	// coloring
	switch ship.name {
	case "WhiteShip":
		ship.setColor("white")
	case "BlackShip":
		ship.setColor("black")
	}

	// notify
	fmt.Printf("%s 로 email 전송 ,%s 를 만들었습니다.", email, ship.getName())
	return ship
}
func main() {
	factory := ShipFactory{}
	whiteShip := factory.orderShip("WhiteShip", "keesun@mail.com")
	fmt.Println(whiteShip.color)
	blackShip := factory.orderShip("BlackShip", "keesun2@mail.com")
	fmt.Println(blackShip.color)
}

//https://refactoring.guru/ko/design-patterns/factory-method/go/example
