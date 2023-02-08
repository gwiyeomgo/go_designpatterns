package main

import "fmt"

type IndoorLighting struct {
	isRunning bool
}

func (t *IndoorLighting) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")

}

func (t *IndoorLighting) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
