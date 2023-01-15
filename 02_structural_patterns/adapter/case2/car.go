package main

import "fmt"

type Car struct {
}

func (c *Car) driveToDestination() {
	fmt.Println("자동차가 목적지로 이동하고 있습니다.")
}
