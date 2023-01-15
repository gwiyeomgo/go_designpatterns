package main

import "fmt"

type Boat struct {
}

func (b *Boat) navigateToDestination() {
	fmt.Println("보트가 부산으로 이동하고 있습니다.")
}
