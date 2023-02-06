package main

import "fmt"

// 작동되거나, 강한 영향을 받고, 무언가에 의해 야기되는 것
type Driven interface {
	drive()
}

type Car struct{}

func (c *Car) drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car   Car
	diver *Driver
}

func (c *CarProxy) drive() {
	if c.diver.Age >= 16 {
		c.car.drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{car: Car{}, diver: driver}
}

func main() {
	car := NewCarProxy(&Driver{Age: 18})
	car.drive()
}
