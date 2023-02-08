package main

import "fmt"

type PokemonGO struct {
	isRunning bool
}

func (p *PokemonGO) start() {
	p.isRunning = true
	fmt.Println("PokemonGO start")

}

func (p *PokemonGO) end() {
	p.isRunning = false
	fmt.Println("PokemonGO end")
}
