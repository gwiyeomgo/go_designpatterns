package case1

//Abstract factory interface
import "fmt"

// Concrete product
// 위임
type IShipFactory interface {
	createShip() IShip
}

func GetShipFactory(name string) (IShipFactory, error) {
	if name == "white" {
		return &WhiteShipFactory{}, nil
	}

	if name == "black" {
		return &BlackShipFactory{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
