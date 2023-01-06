package after

type Ship struct {
	color string
	name  string
}

type IShip interface {
	setColor(name string)
	setName(name string)
	getName() string
	getShip() Ship
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
func (s *Ship) getShip() Ship {
	return *s
}
