package main

import "fmt"

type IAttackStrategy interface {
	attack()
}

type MissileStrategy struct {
}

func (a *MissileStrategy) attack() {
	fmt.Println("I have Missile")
}

type PunchStrategy struct {
}

func (a *PunchStrategy) attack() {
	fmt.Println("I have strong punch")
}
