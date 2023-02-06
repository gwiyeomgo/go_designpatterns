package main

import (
	"fmt"
	"math/rand"
)

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func newCircle(color string) *Circle {
	return &Circle{color: color,
		x:      rand.Intn(100) * 10,
		y:      rand.Intn(100) * 20,
		radius: rand.Intn(100) * 10,
	}
}
func (c *Circle) draw() {
	fmt.Println(fmt.Sprintf("Circle [color= %s , x= %d , y= %d , radius= %d ]", c.color, c.x, c.y, c.radius))
}
