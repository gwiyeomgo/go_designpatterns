package main

import "fmt"

// concrete strategy
type MissileStrategy struct {
}

func (a *MissileStrategy) attack() {
	fmt.Println("I have Missile")
}
