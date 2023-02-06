package main

import "math/rand"

func main() {
	colors := []string{"Red", "Yellow", "Pink", "Blue"}

	for i := 0; i < 10; i++ {
		shapeFactory := getShapeFactorySingleInstance()
		circle := shapeFactory.getCircle(colors[rand.Intn(4)])
		circle.draw()
	}
}
