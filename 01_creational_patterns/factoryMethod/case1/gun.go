package main

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// 어떤 요구사항으로 계속 변경됨
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}
