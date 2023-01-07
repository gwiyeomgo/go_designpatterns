package case2

// 인터페이스 적용하기
type IShipFactory interface {
	order(name string) *Ship
}

func NewFactory(color string) IShipFactory {
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
