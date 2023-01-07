package case1

type WhiteShipFactory struct {
}
type whiteShip struct {
	Ship
}

func (a *WhiteShipFactory) order() IShip {
	ship := whiteShip{
		Ship: Ship{
			logo: "W",
			size: 14,
		},
	}
	return &ship
}
