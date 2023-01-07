package case1

import (
	"fmt"
)

func printShiplDetails(s IShip) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

//어떤 팩토리를 넣느냐에 따라서 제품군의 내용이 달라진다.

func main() {
	whiteShipFactory, _ := GetShipFactory("white")
	blackShipFactory, _ := GetShipFactory("black")

	whiteShipAnchor := whiteShipFactory.order()
	blackShipWheel := blackShipFactory.order()

	printShiplDetails(whiteShipAnchor)
	printShiplDetails(blackShipWheel)
}
