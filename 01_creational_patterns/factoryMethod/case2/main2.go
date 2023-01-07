package case2

import "fmt"

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
