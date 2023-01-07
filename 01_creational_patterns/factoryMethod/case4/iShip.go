package case1

type IShip interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type Ship struct {
	logo string
	size int
}

func (s *Ship) setLogo(logo string) {
	s.logo = logo
}

func (s *Ship) getLogo() string {
	return s.logo
}

func (s *Ship) setSize(size int) {
	s.size = size
}

func (s *Ship) getSize() int {
	return s.size
}
