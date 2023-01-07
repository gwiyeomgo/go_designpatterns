package case1

type BlackShipFactory struct {
}
type blackShip struct {
	Ship
}

func (a *BlackShipFactory) order() IShip {
	return &blackShip{
		Ship: Ship{
			logo: "W",
			size: 14,
		},
	}
}
