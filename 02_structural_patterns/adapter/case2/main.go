package main

import "fmt"

type CarAdapter struct {
	carTransport *Car
}

func (c CarAdapter) navigateToDestination() {
	fmt.Println("내비게이션이 가능하도록 자동차를 개조하다")
	c.carTransport.driveToDestination()
}

func main() {
	client := &Client{}
	boat := &Boat{}

	client.startNavigation(boat)

	car := &Car{}
	carAdapter := &CarAdapter{
		carTransport: car,
	}

	client.startNavigation(carAdapter)
}
