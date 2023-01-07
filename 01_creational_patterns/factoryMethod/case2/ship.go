package main

type Ship struct {
	name  string
	color string
}

func New(name string, color string) *Ship {
	return &Ship{
		name:  name,
		color: color,
	}
}
func (s *Ship) getName() string {
	return s.name
}
